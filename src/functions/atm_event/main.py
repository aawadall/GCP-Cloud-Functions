def atm_event(data, context):
    import base64

    print(data)
    if 'data' in data:
        name = base64.b64decode(data['data'].decode('utf-8'))
    else:
        name = 'Default Name'
    print('Name: {}'.format(name))

