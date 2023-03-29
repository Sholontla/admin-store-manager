from fastapi.staticfiles import StaticFiles
from fastapi.templating import Jinja2Templates
from fastapi import FastAPI, WebSocket, Request, WebSocketDisconnect

from typing import Dict
import json

from sales_trends_analysis import MarketingAnalysis
from descriptive_analytics import DescriptiveAnalyticsSalesByTime
from data_cache import DataCache


class ConnectionManager:
    def __init__(self):
        self.active_connections = []
        self.waiting_connections = []

    async def connect(self, websocket: WebSocket):
        if not self.waiting_connections:
            await websocket.accept()
            self.active_connections.append(websocket)
        else:
            waiting_websocket = self.waiting_connections.pop(0)
            waiting_websocket.__dict__.update(websocket.__dict__)
            await self.connect(waiting_websocket)

    def disconnect(self, websocket: WebSocket):
        self.active_connections.remove(websocket)

    async def broadcast(self, data: Dict):
        for connection in self.active_connections:
            await connection.send_json(data)

    async def wait_for_connection(self):
        websocket = WebSocket(scope={}, receive=None, send=None)
        self.waiting_connections.append(websocket)
        await websocket.receive()
        return websocket


app = FastAPI()
manager = ConnectionManager()


app.mount("/static", StaticFiles(directory="static"), name="static")
templates = Jinja2Templates(directory="templates")


@app.websocket("/ws")
async def websocket_endpoint(websocket: WebSocket):
    await manager.connect(websocket)
    try:
        while True:
            data = await websocket.receive_json()
            cache = DataCache()
            cache.cache_data(data)
            await manager.broadcast(data)
    except WebSocketDisconnect:
        manager.disconnect(websocket)


@app.get("/")
async def root(request: Request):
    return templates.TemplateResponse("index.html", {"request": request})


@app.get("/chartproduct")
async def root(request: Request):
    cache = DataCache()
    all_data_by_time = cache.get_all_data()
    marketing = MarketingAnalysis(all_data_by_time)
    return templates.TemplateResponse("chart_products.html", {"request": request, "sales_data": marketing.sales_by_item()})


@app.get("/products")
async def products():
    global received_data
    marketing = MarketingAnalysis(received_data)
    return marketing.sales_by_item()


@app.get("/chart/sales_month")
async def sales_month(request: Request):
    global received_data
    marketing = MarketingAnalysis(received_data)
    print(marketing.sales_by_month())
    return templates.TemplateResponse("sales_month.html", {"request": request, "sales_data": marketing.sales_by_month()})


@app.get("/sales_month")
async def sales_month():
    global received_data
    marketing = MarketingAnalysis(received_data)
    return marketing.sales_by_month()


@app.get("/test")
async def chart_js(request: Request):
    return templates.TemplateResponse("charts-by-store.html", {"request": request})


@app.get("/tablesdata")
async def tables_data(request: Request):
    global received_data
    marketing = MarketingAnalysis(received_data)
    return templates.TemplateResponse("tables-data.html", {"request": request})


# @app.get("/chartbystore")
# def chartbystore(request: Request):

#     diagnostics = DescriptiveAnalyticsSalesByTime(sales_data)
#     sales_by_store_data = diagnostics.sales_by_store()

#     store_names = list(sales_by_store_data.keys())
#     total_sales = list(sales_by_store_data.values())

#     # Create a dictionary with the data needed to create a Chart.js bar chart
#     chart_data = {
#         "labels": store_names,
#         "datasets": [
#             {
#                 "label": "Total Sales by Store",
#                 "data": total_sales,
#                 "backgroundColor": "rgba(54, 162, 235, 0.2)",
#                 "borderColor": "rgba(54, 162, 235, 1)",
#                 "borderWidth": 1
#             }
#         ]
#     }

#     # Convert the chart_data dictionary to JSON format
#     chart_data_json = json.dumps(chart_data)

#     # Render the HTML template that contains the Chart.js code
#     return templates.TemplateResponse("charts-by-store.html", {"request": request, "chart_data": chart_data_json})


@app.get("/chartbytime")
def chartbystore(request: Request):
    cache = DataCache()
    all_data_by_time = cache.get_all_data()
    diagnostics = DescriptiveAnalyticsSalesByTime(all_data_by_time)
    # Call the sales_by_store() function to retrieve the data
    sales_by_day_store_data = diagnostics.sales_by_day()
    sales_by_week_store_data = diagnostics.sales_by_week()
    sales_by_month_store_data = diagnostics.sales_by_month()
    sales_by_year_store_data = diagnostics.sales_by_year()

    # Extract the store names and total sales from the dictionary
    store_day_names = list(sales_by_day_store_data.keys())
    total_day_sales = list(sales_by_day_store_data.values())

    store_week_names = list(sales_by_week_store_data.keys())
    total_week_sales = list(sales_by_week_store_data.values())

    store_month_names = list(sales_by_month_store_data.keys())
    total_month_sales = list(sales_by_month_store_data.values())

    store_year_names = list(sales_by_year_store_data.keys())
    total_year_sales = list(sales_by_year_store_data.values())

    # Create a dictionary with the data needed to create a Chart.js bar chart
    chart_data = {
        "labels": store_day_names,
        "datasets": [
            {
                "label": "store day",
                "data": total_day_sales,
                "backgroundColor": "rgba(54, 162, 235, 0.2)",
                "borderColor": "rgba(54, 162, 235, 1)",
                "borderWidth": 1
            }
        ]
    }
    chart_week_data = {
        "labels": store_week_names,
        "datasets": [
            {
                "label": 'Stores week',
                "data": total_week_sales,
                "backgroundColor": [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(255, 159, 64, 0.2)',
                    'rgba(255, 205, 86, 0.2)',
                    'rgba(75, 192, 192, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                    'rgba(153, 102, 255, 0.2)',
                    'rgba(201, 203, 207, 0.2)'
                ],
                "borderColor": [
                    'rgb(255, 99, 132)',
                    'rgb(255, 159, 64)',
                    'rgb(255, 205, 86)',
                    'rgb(75, 192, 192)',
                    'rgb(54, 162, 235)',
                    'rgb(153, 102, 255)',
                    'rgb(201, 203, 207)'
                ],
                "borderWidth": 1
            }
        ]
    }

    chart_month_data = {
        "labels": store_month_names,
        "datasets": [
            {
                "label": 'Stores month',
                "data": total_month_sales,
                "backgroundColor": [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(255, 159, 64, 0.2)',
                    'rgba(255, 205, 86, 0.2)',
                    'rgba(75, 192, 192, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                    'rgba(153, 102, 255, 0.2)',
                    'rgba(201, 203, 207, 0.2)'
                ],
                "borderColor": [
                    'rgb(255, 99, 132)',
                    'rgb(255, 159, 64)',
                    'rgb(255, 205, 86)',
                    'rgb(75, 192, 192)',
                    'rgb(54, 162, 235)',
                    'rgb(153, 102, 255)',
                    'rgb(201, 203, 207)'
                ],
                "borderWidth": 1
            }
        ]
    }

    chart_year_data = {
        "labels": store_year_names,
        "datasets": [
            {
                "label": 'Stores year',
                "data": total_year_sales,
                "backgroundColor": [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(255, 159, 64, 0.2)',
                    'rgba(255, 205, 86, 0.2)',
                    'rgba(75, 192, 192, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                    'rgba(153, 102, 255, 0.2)',
                    'rgba(201, 203, 207, 0.2)'
                ],
                "borderColor": [
                    'rgb(255, 99, 132)',
                    'rgb(255, 159, 64)',
                    'rgb(255, 205, 86)',
                    'rgb(75, 192, 192)',
                    'rgb(54, 162, 235)',
                    'rgb(153, 102, 255)',
                    'rgb(201, 203, 207)'
                ],
                "borderWidth": 1
            }
        ]
    }

    # Convert the chart_data dictionary to JSON format
    chart_day_data_json = json.dumps(chart_data)
    chart_week_data_json = json.dumps(chart_week_data)
    chart_month_data_json = json.dumps(chart_month_data)
    chart_year_data_json = json.dumps(chart_year_data)

    # Render the HTML template that contains the Chart.js code
    return templates.TemplateResponse("charts-by-time.html", {"request": request, "chart_data": chart_day_data_json, "chart_week_data": chart_week_data_json, "chart_month_data": chart_month_data_json, "chart_year_data": chart_year_data_json})
