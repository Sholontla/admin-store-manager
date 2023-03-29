import datetime
import string
import websocket
import json
from faker import Faker
import random
import faker_commerce
import time

fake = Faker()
fake.add_provider(faker_commerce.Provider)


ws = websocket.create_connection("ws://localhost:8000/ws")

# define generate_sales_data function on client side


def generate_sales_data(num_sales):
    sales_data = []
    for i in range(num_sales):
        date = fake.date_between(
            start_date='-100d', end_date='today').strftime('%Y-%m-%d')
        store_location = fake.address()
        customer_name = fake.name()
        customer_email = fake.email()
        customer_phone = fake.phone_number()
        items = []
        for j in range(random.randint(1, 5)):

            item_name = fake.ecommerce_name()
            quantity = random.randint(1, 10)
            price = round(random.uniform(0.99, 9.99), 2)
            item_data = {"item_name": item_name,
                         "quantity": quantity, "price": price}
            items.append(item_data)
        total = round(sum(item['quantity'] * item['price']
                      for item in items), 2)
        payment_type = random.choice(["cash", "credit"])
        sale_data = {"date": date, "store_location": store_location, "customer_info": {"name": customer_name,
                                                                                       "email": customer_email, "phone": customer_phone}, "items": items, "total": total, "payment_type": payment_type}
        sales_data.append(sale_data)
    return sales_data


class DataGenerator:
    def __init__(self, store, region, start_date, end_date, product_categories):
        self.store = store
        self.region = region
        self.start_date = datetime.datetime.strptime(
            start_date, '%Y-%m-%d %H:%M:%S')
        self.end_date = datetime.datetime.strptime(
            end_date, '%Y-%m-%d %H:%M:%S')
        self.product_categories = product_categories
        self.counter = 1

    def generate_data(self):
        data = []
        delta = self.end_date - self.start_date
        for i in range(delta.days + 1):
            date = self.start_date + datetime.timedelta(days=i)
            for j in range(random.randint(1, 5)):
                product_category = random.choice(self.product_categories)
                price = random.randint(10, 100)
                quantity = random.randint(1, 20)
                product = {'product_category': product_category,
                           'price': price, 'quantity': quantity}
                products = [product for _ in range(quantity)]
                record = {'id': self.counter, 'store': self.store, 'Region': self.region, 'timestamp': str(date), 'total_sales': price*quantity,
                          'is_promotion': False, 'products': products}
                data.append(record)
                self.counter += 1
        return data


def dummy_generation():
    generator = DataGenerator('Store A', 'A1', '2022-12-19 09:00:00',
                              '2022-12-25 09:00:00', ['Category A', 'Category B', 'Category C'])
    return generator.generate_data()


# send generate_sales_data function with argument to server
data = {"function": "generate_sales_data", "args": [10]}
# d = generate_sales_data(100)
# ws.send(json.dumps(d))

d = dummy_generation()
generator = DataGenerator('Store A', 'A1', '2022-12-19 09:00:00',
                          '2022-12-25 09:00:00', ['Category A', 'Category B', 'Category C'])

data = generator.generate_data()
ws.send(json.dumps(data))

result = ws.recv()
ws.close()


# while True:
#     # send generate_sales_data function with argument to server
#     data = {"function": "generate_sales_data", "args": [10]}
#     d = generate_sales_data(100)
#     ws.send(json.dumps(d))

#     # receive data from server
#     result = ws.recv()
#     print(result)
