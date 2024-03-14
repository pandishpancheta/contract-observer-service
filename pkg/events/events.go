package events

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofrs/uuid/v5"
	"github.com/pandishpancheta/contract-observer-service/pkg/config"
	"github.com/pandishpancheta/contract-observer-service/pkg/protolistings"
	listings "github.com/pandishpancheta/contract-observer-service/pkg/protolistings/pb"
	"github.com/pandishpancheta/contract-observer-service/pkg/protoorders"
	pb "github.com/pandishpancheta/contract-observer-service/pkg/protoorders/pb"
	"github.com/pandishpancheta/contract-observer-service/pkg/wsclient"
)

//event ItemAdded(bytes32 itemId, address owner, string token, uint256 priceInWei);
// event ItemPurchased(bytes32 itemId, address buyer, uint256 tokenId);

type ItemAdded struct {
	ItemId uuid.UUID `json:"itemId"`
}

type ItemPurchased struct {
	OrderId uuid.UUID `json:"orderId"`
}

func SubscribeToEvents(c *ethclient.Client, contractAddress common.Address, logs chan<- types.Log, contractABI abi.ABI) (ethereum.Subscription, error) {

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{contractABI.Events["ItemAdded"].ID}, {contractABI.Events["ItemPurchased"].ID}},
	}

	subscription, err := c.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func DecodeAddEvent(c *ethclient.Client, contractABI abi.ABI, vLog types.Log) (ItemAdded, error) {
	event := contractABI.Events["ItemAdded"]

	addEventMap := make(map[string]interface{})
	err := contractABI.UnpackIntoMap(addEventMap, event.Name, vLog.Data)
	if err != nil {
		return ItemAdded{}, err
	}

	itemIdBytes := addEventMap["orderId"].([32]uint8)
	// itemId is bytes32 of a uuid without dashes
	// convert from bytes32 to string and add dashes to make it a valid uuid
	itemIdFromBytes := string(itemIdBytes[:])

	itemId := uuid.FromStringOrNil(strings.Join([]string{itemIdFromBytes[:8], itemIdFromBytes[8:12], itemIdFromBytes[12:16], itemIdFromBytes[16:20], itemIdFromBytes[20:]}, "-"))

	return ItemAdded{
		ItemId: itemId,
	}, nil
}

func DecodePurchaseEvent(c *ethclient.Client, contractABI abi.ABI, vLog types.Log) (ItemPurchased, error) {
	event := contractABI.Events["ItemPurchased"]

	purchaseEventMap := make(map[string]interface{})
	err := contractABI.UnpackIntoMap(purchaseEventMap, event.Name, vLog.Data)
	if err != nil {
		return ItemPurchased{}, err
	}

	orderIdBytes := purchaseEventMap["orderId"].([32]uint8)
	// orderId is bytes32 of a uuid without dashes
	// convert from bytes32 to string and add dashes to make it a valid uuid
	orderIdFromBytes := string(orderIdBytes[:])

	orderId := uuid.FromStringOrNil(strings.Join([]string{orderIdFromBytes[:8], orderIdFromBytes[8:12], orderIdFromBytes[12:16], orderIdFromBytes[16:20], orderIdFromBytes[20:]}, "-"))

	return ItemPurchased{
		OrderId: orderId,
	}, nil
}

func RunSubscription(cfg *config.Config, contractABIPath string) {
	logs := make(chan types.Log)
	defer close(logs)

	contractAddress := common.HexToAddress(cfg.ContractAddress)

	c, err := wsclient.InitWsClient(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	listingClient, err := protolistings.InitServiceClient(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	orderClient, err := protoorders.InitServiceClient(cfg)

	fileBytes, err := os.ReadFile(contractABIPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	fileContents := string(fileBytes)

	contractAbi, err := abi.JSON(strings.NewReader(fileContents))
	if err != nil {
		log.Fatal(err)
		return
	}

	subscription, err := SubscribeToEvents(c, contractAddress, logs, contractAbi)
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		select {
		case err := <-subscription.Err():
			log.Fatal(err)
			return
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case contractAbi.Events["ItemAdded"].ID.Hex():
				itemAdded, err := DecodeAddEvent(c, contractAbi, vLog)
				if err != nil {
					log.Fatal(err)
					return
				}
				log.Println("Item added: ", itemAdded)
				listingClient.UpdateListingStatus(context.Background(), &listings.UpdateListingStatusRequest{
					Id:     itemAdded.ItemId.String(),
					Status: "confirmed",
				})
			case contractAbi.Events["ItemPurchased"].ID.Hex():
				itemPurchased, err := DecodePurchaseEvent(c, contractAbi, vLog)
				if err != nil {
					log.Fatal(err)
					return
				}
				log.Println("Item purchased: ", itemPurchased)
				orderClient.UpdateStatus(context.Background(), &pb.UpdateStatusRequest{
					Id:     itemPurchased.OrderId.String(),
					Status: "confirmed",
				})

			}
		}
	}
}
