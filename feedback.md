To Do:
<!-- 1. Add a readme -->
<!-- 2. Add an amount column in order table  -->
<!-- 3. Add foreign keys  -->
<!-- 4. Remove userID from orderItems after adding foreign keys -->
<!-- 5. Move the connection.go to database -->
<!-- 6. Rename myInterceptor -->
<!-- 7. Look into database connection (its happening twice) of server.go -->
<!-- low - 8. Try to keep the mail in async manner (use go routines) -->
<!-- 9. Add an "isActive" filter to "getAllAgents" -->
<!-- ***10. Instead of returning the query dircetly, read from the database and return it (if success, else send an error) - DO IT FOR ONE RPC
Example: CreateAgent rpc -->
<!-- 11. Process the errors of utils.CheckError() - done -->
<!-- 12. In sqlMock test, check out "HasExpectationMet" and include it  -->
<!-- 13. Make sure to use Errof when the tests are failing in database_test - done -->
Pros:
1. Good Naming Convention
2. Good organization of code
3. Good Table drived tests
