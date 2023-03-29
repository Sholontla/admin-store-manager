from faker import Faker
import random
import json

fake = Faker()


def generate_sales_data(num_sales):
    sales_data = []
    for i in range(num_sales):
        date = fake.date_between(
            start_date='-30d', end_date='today').strftime('%Y-%m-%d')
        store_location = fake.address()
        customer_name = fake.name()
        customer_email = fake.email()
        customer_phone = fake.phone_number()
        items = []
        for j in range(random.randint(1, 5)):
            item_name = fake.word()
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


if __name__ == '__main__':
    sales_data = generate_sales_data(10)
    print(json.dumps(sales_data, indent=4))
