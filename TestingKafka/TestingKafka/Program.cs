// See https://aka.ms/new-console-template for more information

using MassTransit;
using Microsoft.Extensions.DependencyInjection;

try
{
    var services = new ServiceCollection();

    services.AddMassTransit(x =>
    {
        x.UsingInMemory();
        
        x.AddRider(rider =>
        {
            rider.AddConsumer<KafkaMessageConsumer>();
            rider.AddProducer<KafkaMessage>("topic-name");
            rider.UsingKafka((context, k) =>
            {
                k.Host("0.0.0.0:9092");

                k.TopicEndpoint<KafkaMessage>("topic-name", "consumer-group-name",
                    e =>
                    {
                        e.ConfigureConsumer<KafkaMessageConsumer>(context);
                        e.CreateIfMissing(z => {
                            z.NumPartitions = 1;
                            z.ReplicationFactor = 1;
                        });
                    });
            });
        });
    });

    ServiceProvider provider = services.BuildServiceProvider();

    var busControl = provider.GetRequiredService<IBusControl>();

   await busControl.StartAsync();
// Thread.Sleep(TimeSpan.FromSeconds(5));
    var producer = provider.GetRequiredService<ITopicProducer<KafkaMessage>>();
    do
    {
        string value = await Task.Run(() =>
        {
            Console.WriteLine("Enter text (or quit to exit)");
            Console.Write("> ");
            return Console.ReadLine();
        });

        if ("quit".Equals(value, StringComparison.OrdinalIgnoreCase))
            break;

        await producer.Produce(new
        {
            Text = value
        });
    } while (true);
}

catch (Exception ex)
{
    Console.WriteLine(ex.Message);
}

await Task.Delay(TimeSpan.FromMinutes(60));

class KafkaMessageConsumer :
    IConsumer<KafkaMessage>
{
    public Task Consume(ConsumeContext<KafkaMessage> context)
    {
        Console.WriteLine(context.Message.Text);
        return Task.CompletedTask;
    }
}

public record KafkaMessage
{
    public string Text { get; init; }
}