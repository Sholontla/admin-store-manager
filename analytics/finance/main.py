from finance import Finance
import concurrent.futures


# Sample sales data
sales_data = [{'product': 'A', 'price': 10, 'cost': 5, 'units_sold': 100, 'date': '2023-03-22'},
              {'product': 'B', 'price': 15, 'cost': 8,
                  'units_sold': 200, 'date': '2023-03-22'},
              {'product': 'C', 'price': 20, 'cost': 12,
                  'units_sold': 150, 'date': '2023-03-23'},
              {'product': 'D', 'price': 5, 'cost': 2,
               'units_sold': 300, 'date': '2023-03-23'},
              {'product': 'E', 'price': 30, 'cost': 18,
                  'units_sold': 75, 'date': '2023-03-23'}
              ]


finance = Finance()
# Perform steps 1-5 concurrently and in parallel using ThreadPoolExecutor
with concurrent.futures.ThreadPoolExecutor() as executor:
    future1 = executor.submit(finance.calculate_daily_revenue, sales_data)
    future2 = executor.submit(finance.calculate_average_revenue, sales_data)
    future3 = executor.submit(
        finance.identify_top_selling_products, sales_data)
    future4 = [executor.submit(finance.calculate_profit_margin, sale)
               for sale in sales_data]
    future5 = executor.submit(finance.calculate_daily_cost, sales_data)

daily_revenue = future1.result()
average_revenue = future2.result()
top_selling_products = future3.result()
profit_margin = [future.result() for future in future4]
daily_cost = future5.result()

print("Daily Revenue:", daily_revenue)
print("Average Revenue per Product:", average_revenue)
print("Top Selling Products:", top_selling_products)
print("Profit Margin:", profit_margin)
print("Daily Cost:", daily_cost)


daily_net_profit = finance.calculate_daily_net_profit(
    daily_revenue, daily_cost)
print("Daily Net Profit:", daily_net_profit)
finance = Finance()
average_profit_margin = finance.calculate_average_profit_margin(sales_data)
print("Average Profit Margin:", average_profit_margin)

worst_selling_products = finance.identify_worst_selling_products(sales_data)
print("Worst Selling Products:", worst_selling_products)

highest_revenue_store = finance.identify_highest_revenue_store(sales_data)
print("Store with Highest Revenue:", highest_revenue_store)
