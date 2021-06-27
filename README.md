# CLI Marketplace
Command line in-memory core functionality of a Simplistic Marketplace

### Run
#### Option 1: Download the binary
1. Download the [binary](https://github.com/akshatvn/marketPlace/blob/main/marketPlace)
2. Execute it
```shell script
./marketPlace
```
#### Option 2: Build from source
1. Clone the github repository inside your $GOPATH
```shell script
git clone git@github.com:akshatvn/marketPlace.git
```
2. Execute the run.sh file in the root directory
```shell script
./run.sh
``` 


### Interface
'#' is the prompt for user input.

List of commands available:
```shell script
HELP
REGISTER <username>
CREATE_LISTING <username> <title> <description> <price> <category>
DELETE_LISTING <username> <listing_id>
GET_LISTING <username> <listing_id>
GET_CATEGORY <username> <category> {sort_price|sort_time} {asc|dsc}
GET_TOP_CATEGORY <username>
EXIT

```



### Example usage
```shell script
# REGISTER user1
Success

# CREATE_LISTING user1 'Phone model 8' 'Black color, brand new' 1000 'Electronics'
100001

# GET_LISTING user1 100001
Phone model 8|Black color, brand new|1000|2019-02-22 12:34:56|Electronics|user1

# CREATE_LISTING user1 'Black shoes' 'Training shoes' 100 'Sports'
100002

# REGISTER user2
Success

# REGISTER user2
Error - user already existing

# CREATE_LISTING user2 'T-shirt' 'White color' 20 'Sports'
100003

# GET_LISTING user1 100003
T-shirt|White color|20|2019-02-22 12:34:58|Sports|user2

# GET_CATEGORY user1 'Fashion' sort_time asc
Error - category not found

# GET_CATEGORY user1 'Sports' sort_time dsc
T-shirt|White color|20|2019-02-22 12:34:58|Sports|user2
Black shoes|Training shoes|100|2019-02-22 12:34:57|Sports|user1

# GET_CATEGORY user1 'Sports' sort_price dsc
Black shoes|Training shoes|100|2019-02-22 12:34:57|Sports|user1
T-shirt|White color|20|2019-02-22 12:34:58|Sports|user2

# GET_TOP_CATEGORY user1
Sports

# DELETE_LISTING user1 100003
Error - listing owner mismatch

# DELETE_LISTING user2 100003
Success

# GET_TOP_CATEGORY user2
Sports

# DELETE_LISTING user1 100002
Success

# GET_TOP_CATEGORY user1
Electronics

# GET_TOP_CATEGORY user3
Error - unknown user

```

### Known issues
1. Strings with quotes other than enclosing quotes will result in erroneous behaviour: e.g. "John's Silver Toyota Camry"
2. Categories are case sensitive. 'Bicycles' and 'bicycles' will be treated as different categories
3. GET_CATEGORY does not pad fields in order to align the columns.

### Built with
go 1.13


