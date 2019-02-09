/**
 * Triggered from a message on a Cloud Pub/Sub topic.
 *
 * @param {!Object} event Event payload.
 * @param {!Object} context Metadata for the event.
 */
exports.helloPubSub = (event, context) => {
    const pubsubMessage = event.data;
    const pubsubAttributes = event.attributes;

    console.log("New Event=====")
    console.log(Buffer.from(pubsubMessage, 'base64').toString());
    console.log("Attributes:")
    pubsubAttributes.forEach( a => {
        console.log(Buffer.from(a, 'base64').toString());
    });
};
