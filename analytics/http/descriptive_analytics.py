import collections
# Descriptive analytics: Use descriptive analytics techniques to gain insights into sales trends and patterns. This includes analyzing sales by day, week, month, or year to identify trends and patterns. You can also analyze sales by product category, product type, and store location to understand what products are selling well and where.

from datetime import datetime


class DescriptiveAnalyticsSalesByTime:
    def __init__(self, sales_data):
        self.sales_data = sales_data

    def sales_by_day(self):
        sales_by_time_day = {datetime.fromisoformat(sale['timestamp']).strftime(
            '%Y-%m-%d'): sale['total_sales'] for sale in self.sales_data}
        sorted_sales = sorted(
            sales_by_time_day.items(), key=lambda x: datetime.fromisoformat(x[0]))
        sales_by_time_sorted = collections.OrderedDict(sorted_sales)
        keys = sales_by_time_sorted.keys()
        values = sales_by_time_sorted.keys()
        return sales_by_time_sorted

    def sales_by_week(self):
        sales_by_time_week = {datetime.fromisoformat(sale['timestamp']).strftime(
            '%Y-W%U'): sale['total_sales'] for sale in self.sales_data}
        sales_by_time_week = {key: sum(sale['total_sales'] for sale in self.sales_data if datetime.fromisoformat(
            sale['timestamp']).strftime('%Y-W%U') == key) for key in sales_by_time_week}
        return sales_by_time_week

    def sales_by_month(self):
        sales_by_time_month = {datetime.fromisoformat(sale['timestamp']).strftime(
            '%Y-%m'): sale['total_sales'] for sale in self.sales_data}
        sales_by_time_month = {key: sum(sale['total_sales'] for sale in self.sales_data if datetime.fromisoformat(
            sale['timestamp']).strftime('%Y-%m') == key) for key in sales_by_time_month}
        return sales_by_time_month

    def sales_by_year(self):
        sales_by_time_year = {datetime.fromisoformat(sale['timestamp']).strftime(
            '%Y'): sale['total_sales'] for sale in self.sales_data}
        sales_by_time_year = {key: sum(sale['total_sales'] for sale in self.sales_data if datetime.fromisoformat(
            sale['timestamp']).strftime('%Y') == key) for key in sales_by_time_year}
        return sales_by_time_year

    def sales_by_category(self):
        """
        Returns a dictionary with the total sales for each product category.
        """
        sales_by_category = {}
        for sale in self.sales_data:
            for product in sale['products']:
                category = product['product_category']
                sales_by_category[category] = sales_by_category.get(
                    category, 0) + sale['total_sales']
        return sales_by_category

    def sales_by_store(self):
        """
        Returns a dictionary with the total sales for each store location.
        """
        sales_by_store = {}
        for sale in self.sales_data:
            store = sale['store']
            sales_by_store[store] = sales_by_store.get(
                store, 0) + sale['total_sales']
        return sales_by_store


sales_data = [
    {'id': 1, 'store': 'Store A', 'timestamp': '2022-12-19 09:00:00', 'total_sales': 1000, 'is_promotion': False,
        'products': [{'product_category': 'Category A', 'price': 50, 'quantity': 10}]}
]
