def atm_event(data, context):
    import base64

    print(base64.b64decode(data))
    if 'data' in data:
        name = base64.b64decode(data['data'])
    else:
        name = 'Default Name'
    print('Name: {}'.format(name))

