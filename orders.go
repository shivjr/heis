package order

import ("fmt")

type (
Matrix 	= [][]int
Array 	= []int
)








var (
	Direction int
	ExternalQueue [][]int
	InternalQueue []int
	UpdateLightCh = make(chan string, 5)
	//osChan chan os.Signal
)

func Run(){
	ExternalQueue, InternalQueue := initilizeExternalAndInternalQueue()
	Direction := Down

	go AddExternalOrder()

}


AddExternalOrder(){
	var data Message
	var receiverID int
	order := make(int, 2)
	for (){
		data = <- network.AddOrderCh
		order = data.Order
		receiverID = data.ReceiverID		
		ExternalQueue[order[0]][order[1]] = receiver.ID
		fmt.Println("ExternalQueue updated. Now:")
		printExternalQueue()
		
		UpdateLightCh <- "external"
		netWork.OutputCh <- Message{Head: "Queue", Order: order, Type: ExternalQueue}
		

		}
		 
}


////////////////////////////////////////////////////////////////////////////////////
//Help functions
initializeExternalAndInternalQueue() Matrix Array
{
	
}
