






curl -X POST http://localhost:8090/api/collections/users/records \
  -H "Content-Type: application/json" \
  -d '{"email": "john@example.com", "password": "mypassword123", "passwordConfirm": "mypassword123", "name": "John Doe"}'
