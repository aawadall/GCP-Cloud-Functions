/**
 * Responds to any HTTP request.
 *
 * @param {!express:Request} req HTTP request context.
 * @param {!express:Response} res HTTP response context.
 */
exports.helloWorld = (req, res) => {
    if(req.body.message === undefined) {
        console.log('No message defined');
        res.status(400).send('No Message Defined!');
    }  else {
        console.log(req.body.message);
        let message = req.query.message || req.body.message || 'Hello World!';
        res.status(200).send(message);
    }
};
