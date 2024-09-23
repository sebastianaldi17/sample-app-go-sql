# Sample-Todo-App
A sample to-do app. Consists of backend and frontend.

- backend-go: backend written in Go, using chi for HTTP services.
- backend-spring: backend written in Java Spring Boot, using Spring Web starter 
- frontend-vue: frontend written in Vue (typescript), using Vuetify component framework

# To-do (from most interested to least interested)
- Add redis caching (render.com has free redis for experimenting)
- Add logging & tracing to NewRelic
- Try out writer & reader DB separation for PostgreSQL (have a different instance for read + write and read only)
    - Try out reader replication (having more than one instance for reader)
- Implement local CI/CD using Jenkins
- Create backend-typescript
- Create frontend-react
- Add tests 
    - For backend-go, need to change init from direct struct to interface so that mocks can be generated
    - For backend-java, use H2 in memory database, watch this [video](https://youtu.be/Nv2DERaMx-4?t=8485&si=htafZ2XGGj6qTrvI)