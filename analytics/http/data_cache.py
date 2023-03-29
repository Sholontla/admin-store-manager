import redis
import json


class DataCache:
    def __init__(self, host='localhost', port=6379, db=0):
        self.redis = redis.Redis(host=host, port=port,
                                 db=db, password='eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81')

    def cache_data(self, data):
        for record in data:
            # Convert dict to JSON string and set it as a value in Redis with a key equal to the record ID
            record_id = record.get('id')
            if record_id:
                self.redis.set(record_id, json.dumps(record))

    def get_data(self, record_ids):
        data = []
        for record_id in record_ids:
            # Get JSON string from Redis and convert it to a dict
            record_data = self.redis.get(record_id)
            if record_data:
                data.append(json.loads(record_data))
        return data

    def get_all_data(self):
        keys = self.redis.keys()
        data = self.redis.mget(keys)
        return [json.loads(record_data) for record_data in data if record_data is not None]

    def delete_all_data(self):
        self.redis.flushall()

# if __name__ == '__main__':
#     cache = DataCache()

#     # Cache data
#     data = {'id': '3', 'name': 'John'}, {'id': '4', 'name': 'Jane'}
#     cache.cache_data(data)

#     # Get data
#     record_ids = ['3', '2']
#     result = cache.get_data(record_ids)
#     print(result)
