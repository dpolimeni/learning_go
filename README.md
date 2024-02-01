# Create a simple APP with fastapi structure
Here i try to create a simple go application with a fastapi structure referencing to this repo https://github.com/zhanymkanov/fastapi-best-practices.

## Main routes
I define a route with CRUD interaction with a sql DB about people.
Then a route interacting with a NoSQL db... (let's see which one).

### USeful commands
- export PATH=$(go env GOPATH)/bin:$PATH to add swag on your path
- TO GET VISUALIZATION TOOL of your DB: curl -sSf https://atlasgo.sh | sh 
- atlas schema inspect \
    -u "ent://ent/schema" \
    --dev-url "sqlite://file?mode=memory&_fk=1" \ (TO VISUALIZE)
    -w

