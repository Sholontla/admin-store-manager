import json


class DataIngestion:
    def __init__(self, data_source):
        self.data_source = data_source

    def read_data(self):
        if self.data_source.endswith('.csv'):
            with open(self.data_source, 'r') as f:
                data = f.readlines()
        elif self.data_source.endswith('.json'):
            with open(self.data_source, 'r') as f:
                data = json.load(f)
        else:
            raise ValueError('Unsupported data source')

        # Clean the data
        cleaned_data = []
        for row in data:
            cleaned_row = row.strip().split(',')
            cleaned_data.append(cleaned_row)

        return cleaned_data


class DataAnalysis:
    def __init__(self, data):
        self.data = data

    def calculate_metrics(self):
        pass

    def generate_reports(self):
        pass

    def visualize_data(self):
        pass


class PipelineManager:
    def __init__(self, data_sources):
        self.data_sources = data_sources

    def run_pipeline(self):
        for data_source in self.data_sources:
            with DataIngestion(data_source) as di:
                data = di.read_data()

            with DataAnalysis(data) as da:
                da.calculate_metrics()
                da.generate_reports()
                da.visualize_data()


if __name__ == '__main__':
    pipeline_manager = PipelineManager(['data.csv', 'data.json'])
    pipeline_manager.run_pipeline()
