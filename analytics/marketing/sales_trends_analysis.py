import matplotlib.pyplot as plt
import datetime


from collections import defaultdict


class MarketingAnalysis:
    def __init__(self, sales_data):
        self.sales_data = sales_data

    def sales_by_item(self):
        sales_by_item = defaultdict(lambda: {"quantity": 0, "revenue": 0})
        for sale in self.sales_data:
            for item in sale["items"]:
                item_name = item["item_name"]
                sales_by_item[item_name]["quantity"] += item["quantity"]
                sales_by_item[item_name]["revenue"] += item["quantity"] * \
                    item["price"]
        return sales_by_item

    def sales_by_month(self):
        sales_by_month = {}
        for sale in self.sales_data:
            month_year = sale["sale_date"].strftime("%Y-%m")
            if month_year not in sales_by_month:
                sales_by_month[month_year] = {
                    "quantity": 0,
                    "revenue": 0
                }
            sales_by_month[month_year]["quantity"] += sum(
                item["quantity"] for item in sale["items"])
            sales_by_month[month_year]["revenue"] += sale["total"]
        return sales_by_month

    def sales_by_day(self):
        sales_by_day = {}
        today = datetime.date.today()
        for sale in self.sales_data:
            sale_date = sale["sale_date"].date()
            if sale_date == today:
                for item in sale["items"]:
                    item_name = item["item_name"]
                    if item_name not in sales_by_day:
                        sales_by_day[item_name] = 0
                    sales_by_day[item_name] += item["quantity"] * item["price"]
        return sales_by_day

    def sales_by_week(self):
        sales_by_week = {}
        for sale in self.sales_data:
            sale_date = sale["sale_date"]
            year, week, _ = sale_date.isocalendar()
            week_start = sale_date - \
                datetime.timedelta(days=sale_date.weekday())
            week_end = week_start + datetime.timedelta(days=6)
            week_label = f"{year}-W{week}"
            if week_label not in sales_by_week:
                sales_by_week[week_label] = {
                    "quantity": 0,
                    "revenue": 0
                }
            sales_by_week[week_label]["quantity"] += sum(
                item["quantity"] for item in sale["items"])
            sales_by_week[week_label]["revenue"] += sale["total"]
        return sales_by_week

    def plot_today_sales(self):
        sales_by_day = self.sales_by_day()
        item_names = list(sales_by_day.keys())
        revenue = list(sales_by_day.values())
        plt.bar(item_names, revenue)
        plt.xlabel('Item')
        plt.ylabel('Revenue')
        plt.title('Today Sales')
        plt.show()
