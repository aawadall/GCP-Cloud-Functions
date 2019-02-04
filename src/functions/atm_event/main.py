def atm_event(data, context):
    import base64

    print(base64.b64decode(data).decode('utf-8'))
    if 'data' in data:
        name = base64.b64decode(data['data'].decode('utf-8'))
    else:
        name = 'Default Name'
    print('Name: {}'.format(name))

