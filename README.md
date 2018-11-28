# GolangRestfulAPI
Restful API example written in golang

Start the RESTFul Api that runs at port 8000 of localhost execute the command line:

  go run absolute_path_project_folder\main.go

  
To perform requests to the Api you can make the following requests using postman or a browser:

    - http://localhost:8000/people
      Returns:
      [{"ID":"1","Firstname":"Timmy","lastname":"Kerouac","address":{"city":"City X","state":"State X","phone":"+35387388888"}},
	     {"ID":"2","Firstname":"Albert","lastname":"Einstein","address":{"city":"City Z","state":"State Y","phone":"+35385993751"}}]
    
    - http://localhost:8000/people/1
	    Returns:  {"ID":"1","Firstname":"Timmy","lastname":"Kerouac","address":{"city":"City X","state":"State X","phone":"+35387388888"}}
      
    - using Postman or some other utility that allows to specify the POST or DELETE verb
	    
      /people/{id} Methods POST
      
      /people/{id} Methods DELETE

