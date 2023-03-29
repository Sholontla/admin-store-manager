import asyncio
from kafka import KafkaConsumer


async def consume_messages():
    # Create a Kafka consumer client
    consumer = KafkaConsumer(
        'test_topic',  # Topic to consume messages from
        bootstrap_servers=['localhost:9092'],
        auto_offset_reset='earliest',  # Start reading from the beginning of the topic
        enable_auto_commit=True,  # Automatically commit offsets after consuming messages
        group_id='my-group')  # Consumer group ID

    # Continuously consume messages from the Kafka topic
    while True:
        message = await asyncio.get_event_loop().run_in_executor(None, next, consumer)
        print(message.value)


async def main():
    consumer_task = asyncio.create_task(consume_messages())
    await asyncio.gather(consumer_task)

if __name__ == '__main__':
    asyncio.run(main())
