# from __future__ import print_function

# import boto3
# import time

# def lambda_handler(event, context):
#     path = "/"
#     for items in event["Records"]:
#         if path in items["s3"]["object"]["key"]:
#             client = boto3.client('cloudfront')
#             invalidation = client.create_invalidation(DistributionId='EXVC588OTR2KI',
#                 InvalidationBatch={
#                     'Paths': {
#                         'Quantity': 1,
#                         'Items': ['/*']
#                 },
#                 'CallerReference': str(time.time())
#             })
#             break