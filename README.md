### Basic TODO service

### Service for management todo list via API using memory storage. 


#### 1. Create a new task
```bash
curl --location --request POST 'localhost:3000/task' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "test task",
    "description": "test desc"
}'
```

#### 2. Update exists task
```bash
curl --location --request PUT 'localhost:3001/task' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "test task 2",
    "description": "test desc 2 f",
    "id": "2a8febe8-3691-11ee-ad53-def1c19567d2",
    "status": true
}'
```
#### 3. Delete exists task
```bash
curl --location --request DELETE 'localhost:3000/task' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "fd0261dc-3691-11ee-a92f-def1c19567d2"
}'
```
#### 4. Get list registered tasks
```bash
curl --location --request GET 'localhost:3000/list'
```
# Run
```bash
make install 
make start
```