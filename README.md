# **URL Shortener**

Base URL: `127.0.0.1:8000`

| Endpoints            | Method                 |
| -------------------- | ---------------------- |
| `/v1/users`          | `GET` `POST`           |
| `/v1/users/:id`      | `GET` `PATCH` `DELETE` |
| `/v1/users/:id/urls` | `GET`                  |
| `/v1/urls`           | `GET` `POST`           |
| `/v1/urls/:id`       | `GET` `PATCH` `DELETE` |
| `/:shortened-url`    | `GET`                  |

### **How to run**

- Using docker-compose
  ```
  docker compose up
  ```
- Run manually (requires docker)

  1. Clone this repository **OR** copy the `init.sql` to current directory

     ```
     git clone https://github.com/hisyamsk/url-shortener.git
     ```

  2. Create custom docker volume and Run PostgreSQL image

     ```
     docker volume create YOUR_VOLUME
     ```

     ```
       docker run -d \
         --name db \
         -v YOUR_VOLUME:/var/lib/postgresql/data \
         -v "$(pwd)/init.sql:/docker-entrypoint-initdb.d/init.sql" \
         -e POSTGRES_PASSWORD=foobarbaz \
         -p 5432:5432 \
         --restart unless-stopped \
         postgres:15.1-alpine
     ```

  3. Run the app image

     ```
       docker run -d \
         --name univ-api \
         -e DB_USER=postgres \
         -e DB_PASSWORD=foobarbaz \
         -e DB_HOST=db \
         -e DB_PORT=5432 \
         -e APP_PORT=:8000 \
         --link=db \
         -p 8000:8000 \
         --restart unless-stopped \
         hisyamsk/url-shortener

     ```
