### BUILD image
ARG JAVA_VERSION=14
FROM maven:3-jdk-$JAVA_VERSION as builder
WORKDIR /build
COPY pom.xml .
COPY src src
RUN mvn clean install

FROM openjdk:$JAVA_VERSION-slim as runtime
EXPOSE 8080
#Set app home folder
ENV APP_HOME /app

ENV DATABASE_TYPE=h2
ENV DATABASE_URL="mem:tesla-db"

#Create base app folder
WORKDIR $APP_HOME

#Copy executable jar file from the builder image
COPY --from=builder /build/target/demo-java.jar app.jar
COPY entrypoint.sh .
ENTRYPOINT [ "./entrypoint.sh"]
