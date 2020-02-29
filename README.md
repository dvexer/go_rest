# Simple REST server

Usage example:
1) To get all contacts:
curl http://localhost:8000/phones/
2) To get contact Merry:
curl http://localhost:8000/phones/Merry
3) To add contact Merry with phone number 123:
curl -X POST http://localhost:8000/phones -H "Content-Type: application/json" -d "{\"name\": \"Merry\", \"phone\": \"123\"}"
4) To update contact Merry to Ann:
curl -X PUT http://localhost:8000/phones/Merry -H "Content-Type: application/json" -d "{\"name\": \"Ann\", \"phone\": \"123\"}"
5) To delete Merry contact:
curl -X DELETE http://localhost:8000/phones/Merry
