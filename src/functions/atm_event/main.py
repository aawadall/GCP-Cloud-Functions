import base64

def atm_event(data, context):
    print("Message received")
    if 'data' in data:
        name = "Data"
    else:
        name = 'Default Name'
    print('Name: {}'.format(name))

