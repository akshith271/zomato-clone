# Zomato Clone
This is a clone of a food delivery system. This projects helps us in understanding various things like 
**gRPCs**, **Writing tests**, **Data Modelling**, etc 

## Setup Instructions

 -  Download go from [Go website](https://go.dev/doc/install)
 -  To clone the git repository into your local, follow the command below
	> `git clone https://github.com/akshith271/zomato-clone.git`
 - Navigate to the cloned project 
 - Run the command to import the various packages needed :
	> `go mod tidy`
 -  Start the server using  :
	> `go run mock-grpc/server/server.go`
 -  Start the client using :
	> `go run mock-grpc/client/client.go`
- To test the files in the project :
		> `go test -v/ go test` 
		
 - The gRPCs can also be   tested using [Postman](https://www.postman.com/) or [BloomRPC](https://github.com/bloomrpc/bloomrpc)
---
## Database Design & Schema
**![](https://lh3.googleusercontent.com/ZVrt4awRIqPxKdRMmf55Ya20jKGvHJDA2uHBWc50DWDWzIv5DtstfF9hMqjad-1fxkas1emwyLuqgNSGKnKVCCvSYJCrReQKhHhngcmEQ9NykWtHAA-caJuUW6LD-aSDdIO0ZUQ1j-69FsvSrAJHBhi2SfVyxlXMmH97KuW_SJgZgSEqGNOJTKwmiuEkiQ)**
## Use Cases
 
### Users can :
    

- Create their profile
    
- Get their profile - view their profile
    
- Update their profile
    
- Get the food menu based on restaurant
    
- Place an order (assuming payment was successful)
    
- View their previous order items
### Restaurants can:
    

- Register in Zomato
    
- Can add dishes to their menu
    
- Update their restaurant profile
    

### Delivery Agents can:
    

- Register themselves
    

- Update their status to active/ inactive

---
