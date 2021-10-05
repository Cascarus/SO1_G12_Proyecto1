// Imports the Google Cloud client library
import { PubSub } from '@google-cloud/pubsub'


export default async function subscribeWithFlowControlSettings(maxInProgress, subscriptionName, timeout) {
    
    const pubSubClient = new PubSub();

    const subscriberOptions = {
        flowControl: {
            maxMessages: maxInProgress,
        },
    };

    // References an existing subscription.
    // Note that flow control settings are not persistent across subscribers.
    const subscription = pubSubClient.subscription(
        subscriptionName,
        subscriberOptions
    );

    console.log(
        `Subscriber to subscription ${subscription.name} is ready to receive messages at a controlled volume of ${maxInProgress} messages.`
    );

    const messageHandler = message => {
        console.log(`Received message: ${message.id}`);
        console.log(`\tData: ${message.data}`);
        console.log(`\tAttributes: ${JSON.stringify(message.attributes)}`);


        // "Ack" (acknowledge receipt of) the message
        message.ack();
    };

    subscription.on('message', messageHandler);

    setTimeout(() => {
        console.log("Cerrando suscripcion...")
        subscription.close();
    }, timeout * 1000);

}
