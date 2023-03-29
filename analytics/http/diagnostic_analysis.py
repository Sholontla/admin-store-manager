
# Descriptive analytics: Use descriptive analytics techniques to gain insights into sales trends and patterns. This includes analyzing sales by day, week, month, or year to identify trends and patterns. You can also analyze sales by product category, product type, and store location to understand what products are selling well and where.
from datetime import datetime


class DiagnosticAnalytics:
    def __init__(self, sales_data):
        self.sales_data = sales_data

    def analyze_seasonality(self, time_period='month'):
        """
        Analyze sales data for seasonality based on a given time period
        (e.g. month, quarter, year).

        Returns a dictionary with the total sales for each time period.
        """
        sales_by_period = {}
        for sale in self.sales_data:
            timestamp = sale['timestamp']
            timestamp = datetime.strptime(timestamp, '%Y-%m-%d %H:%M:%S')
            time_period_val = None
            if time_period == 'month':
                time_period_val = timestamp.strftime('%Y/%m')
            elif time_period == 'quarter':
                quarter_num = (timestamp.month - 1) // 3 + 1
                time_period_val = f"{timestamp.year} Q{quarter_num}"
            elif time_period == 'year':
                time_period_val = timestamp.year

            if time_period_val in sales_by_period:
                sales_by_period[time_period_val] += sale['total_sales']
            else:
                sales_by_period[time_period_val] = sale['total_sales']

        return sales_by_period

    def analyze_promotions(self):
        """
        Analyze the impact of promotions on sales.

        Returns a dictionary with the total sales during and outside of promotions.
        """
        promo_sales = 0
        non_promo_sales = 0
        for sale in self.sales_data:
            if sale['is_promotion']:
                promo_sales += sale['total_sales']
            else:
                non_promo_sales += sale['total_sales']

        return {
            'promo_sales': promo_sales,
            'non_promo_sales': non_promo_sales
        }

    def analyze_pricing(self):
        """
        Analyze the impact of pricing on sales.

        Returns a dictionary with the average price and total sales for each product category.
        """
        sales_by_category = {}
        for sale in self.sales_data:
            for product in sale['products']:
                category = product['product_category']
                price = product['price']
                quantity = product['quantity']
                sales = price * quantity

                if category in sales_by_category:
                    sales_by_category[category]['total_sales'] += sales
                    sales_by_category[category]['total_quantity'] += quantity
                else:
                    sales_by_category[category] = {
                        'total_sales': sales,
                        'total_quantity': quantity
                    }

        for category, sales_info in sales_by_category.items():
            sales_info['avg_price'] = sales_info['total_sales'] / \
                sales_info['total_quantity']

        return sales_by_category

    def identify_outliers(self, threshold=3):
        """
        Identify outliers in the sales data using the Z-score method.

        Returns a list of sales that are outliers.
        """
        sales = [sale['total_sales'] for sale in self.sales_data]
        mean_sales = sum(sales) / len(sales)
        std_dev = (
            sum([(sale - mean_sales) ** 2 for sale in sales]) / len(sales)) ** 0.5

        outliers = []
        for i, sale in enumerate(self.sales_data):
            z_score = (sale['total_sales'] - mean_sales) / std_dev
            if abs(z_score) > threshold:
                outliers.append(sale)

        return outliers

    def analyze(self):
        """
        Perform all diagnostic analytics techniques and return the results as a dictionary.
        """
        results = {}
        results['seasonality'] = self.analyze_seasonality()
        results['promotions'] = self.analyze_promotions()
        results['pricing'] = self.analyze_pricing()
        results['outliers'] = self.identify_outliers()

        return results


sales_data = [
    {'store': 'Store A', 'timestamp': '2022-01-01 09:00:00', 'total_sales': 1000, 'is_promotion': False,
        'products': [{'product_category': 'Category A', 'price': 50, 'quantity': 10}]},
    {'store': 'Store A', 'timestamp': '2022-01-01 12:00:00', 'total_sales': 1500, 'is_promotion': True,
        'products': [{'product_category': 'Category B', 'price': 75, 'quantity': 10}]},
    {'store': 'Store B', 'timestamp': '2022-01-01 15:00:00', 'total_sales': 2000, 'is_promotion': False,
        'products': [{'product_category': 'Category A', 'price': 100, 'quantity': 10}, {'product_category': 'Category B', 'price': 100, 'quantity': 10}]},
    {'store': 'Store C', 'timestamp': '2022-01-02 10:00:00', 'total_sales': 3000, 'is_promotion': False,
        'products': [{'product_category': 'Category B', 'price': 150, 'quantity': 10}]},
    {'store': 'Store A', 'timestamp': '2022-01-02 14:00:00', 'total_sales': 2500, 'is_promotion': True,
        'products': [{'product_category': 'Category A', 'price': 50, 'quantity': 25}, {'product_category': 'Category B', 'price': 100, 'quantity': 10}]},
    {'store': 'Store B', 'timestamp': '2022-01-03 11:00:00', 'total_sales': 1000, 'is_promotion': False,
        'products': [{'product_category': 'Category A', 'price': 20, 'quantity': 10}]},
    {'store': 'Store C', 'timestamp': '2022-01-03 14:00:00', 'total_sales': 5000, 'is_promotion': True,
        'products': [{'product_category': 'Category A', 'price': 100, 'quantity': 20}, {'product_category': 'Category B', 'price': 200, 'quantity': 10}]},
]

if __name__ == '__main__':
    diagnostics = DiagnosticAnalytics(sales_data)

    print(diagnostics.analyze())
