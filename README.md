# Kafka Message Scheduler

## Overview

This repository contains a Kafka message scheduler application designed with the aim of efficiently scheduling millions of messages. The application allows you to schedule and produce messages to a Kafka topic. It utilizes Docker Compose to set up a Kafka cluster, PostgreSQL database, and Debezium connector. Follow the instructions below to get started.

## Prerequisites

Make sure you have Docker and Docker Compose installed on your machine.

## Getting Started

1. Clone this repository:

   ```bash
   git clone https://github.com/devharshgupta/kafka-message-scheduler.git
   ```

2. Navigate to the project directory:

   ```bash
   cd kafka-message-scheduler
   ```

3. Start the Kafka cluster, PostgreSQL, and Debezium:

   ```bash
   docker-compose up
   ```

4. After the cluster starts successfully, run the application:

   ```bash
   go run main.go
   ```

## Debezium Setup

To enable Debezium to listen to changes in PostgreSQL, use the following CURL command:

```bash
curl --location 'http://localhost:8083/connectors' \
--header 'Content-Type: application/json' \
--data '{
  "name": "scheduled-messages",
  "config": {
    "connector.class": "io.debezium.connector.postgresql.PostgresConnector",
    "plugin.name": "pgoutput",
    "database.hostname": "host.docker.internal",
    "database.port": "5432",
    "database.user": "docker",
    "database.password": "docker",
    "database.dbname": "messages",
    "database.server.name": "messages",
    "table.include.list": "public.messages",
    "topic.prefix": "scheduled-messages",
    "transforms": "unwrap",
    "transforms.unwrap.type": "io.debezium.transforms.ExtractNewRecordState",
    "transforms.unwrap.drop.tombstones": "false",
    "transforms.unwrap.delete.handling.mode": "rewrite",
    "transforms.unwrap.operation.header": "true",
    "key.converter": "org.apache.kafka.connect.json.JsonConverter",
    "value.converter": "org.apache.kafka.connect.json.JsonConverter",
    "value.converter.schemas.enable": "false",
    "key.converter.schemas.enable": "false"
  }
}'
```

## Health Check

Check the health of the application with the following CURL command:

```bash
curl --location 'localhost:3000'
```

## Message Priority Levels

The application supports three priority levels for scheduling tasks:

- **Priority -1 (P-1):** Immediate tasks that need to be scheduled without any delay.

- **Priority 0 (P0):** Tasks that need to be scheduled after some time.

- **Priority 1 (P1):** Tasks that can be scheduled, for example, an hour from now, as they are intended for a later point in time.

## Producing a Million Messages

To produce millions of messages, feel free to customize the application and leverage its scalability. Use the provided CURL command as a starting point:

```bash
curl --location 'localhost:3000/v1/scheduler/message' \
--header 'Content-Type: application/json' \
--data '{
    "Priority" : -1,
    "Key": "Hi",
    "Value": {
        "message" : "Hello Harsh Gupta"
    },
    "ScheduledAt": "2024-02-01T01:59:00Z"
}'
```

## Connect for Further Discussion

If you have any questions, suggestions, or would like to discuss collaboration opportunities, feel free to connect with me on [LinkedIn](https://www.linkedin.com/in/devharshgupta/). I am more than happy to engage in discussions and explore potential collaborations on this project or any other ideas you may have. Looking forward to connecting with you!