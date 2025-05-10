![alt text](./docs/image.png)

## Database
```bash
docker run --name my-postgres \
  -e POSTGRES_USER=admin \
  -e POSTGRES_PASSWORD=secret \
  -e POSTGRES_DB=studentdb \
  -p 5432:5432 \
  -v $(pwd)/pg_data:/var/lib/postgresql/data \
  -d postgres
```

- Tạo rest api server
    - Thêm sửa xoá students
    - Kết nối với postgresql
- Start postgresql