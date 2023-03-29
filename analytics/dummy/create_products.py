import json
from contextlib import contextmanager
from faker import Faker
import faker_commerce


@contextmanager
def save_file(filename):
    try:
        file = open(filename, 'w')
        yield file
    finally:
        file.close()


def products_random_semect():
    fake = Faker()
    fake.add_provider(faker_commerce.Provider)

    n = 1000  # number of products to generate
    products = []
    for _ in range(n):
        product_name = fake.ecommerce_name()
        products.append(product_name)

    s = set(products)

    # Example usage
    with save_file('products_list.json') as f:
        f.write(json.dumps({"products": list(s)}))


if __name__ == '__main__':
    products_random_semect()
