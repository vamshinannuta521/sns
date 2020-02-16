cd /opt
wget http://www-us.apache.org/dist/kafka/2.1.1/kafka_2.11-2.1.1.tgz
tar xvzf kafka_2.11-2.1.1.tgz
ln -s kafka_2.11-2.1.1/ kafka


vi /etc/systemd/system/kafka-zookeeper.service
[Unit]
Description=Apache Zookeeper server (Kafka)
Documentation=http://zookeeper.apache.org
Requires=network.target remote-fs.target
After=network.target remote-fs.target

[Service]
Type=simple
User=root
Group=root
Environment=JAVA_HOME=/usr/java/jdk1.8.0_102
ExecStart=/opt/kafka/bin/zookeeper-server-start.sh /opt/kafka/config/zookeeper.properties
ExecStop=/opt/kafka/bin/zookeeper-server-stop.sh

[Install]
WantedBy=multi-user.target


vi /etc/systemd/system/kafka.service
[Unit]
Description=Apache Kafka server (broker)
Documentation=http://kafka.apache.org/documentation.html
Requires=network.target remote-fs.target
After=network.target remote-fs.target kafka-zookeeper.service

[Service]
Type=simple
User=root
Group=root
Environment=JAVA_HOME=/usr/java/jdk1.8.0_102
ExecStart=/opt/kafka/bin/kafka-server-start.sh /opt/kafka/config/server.properties
ExecStop=/opt/kafka/bin/kafka-server-stop.sh

[Install]
WantedBy=multi-user.target


vi kafka/config/server.properties
edit listeners propertie
listeners=PLAINTEXT://localhost:9092


systemctl start kafka-zookeeper.service
systemctl start kafka.service
