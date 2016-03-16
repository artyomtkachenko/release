
# Disable the stupid stuff rpm distros include in the build process by default:
#   Disable any prep shell actions. replace them with simply 'true'
%define __spec_prep_post true
%define __spec_prep_pre true
#   Disable any build shell actions. replace them with simply 'true'
%define __spec_build_post true
%define __spec_build_pre true
#   Disable any install shell actions. replace them with simply 'true'
%define __spec_install_post true
%define __spec_install_pre true
#   Disable any clean shell actions. replace them with simply 'true'
%define __spec_clean_post true
%define __spec_clean_pre true
# Disable checking for unpackaged files ?
#%undefine __check_files

# Use md5 file digest method
%define _binary_filedigest_algorithm 1

# Use gzip payload compression
%define _binary_payload w9.gzdio 

Name: foo
Version: 1.0.1        
Release:        1
Summary: blah        
AutoReqProv: no
BuildRoot: /home/home/git/release/data/foo4754965266/BUILD
# Seems specifying BuildRoot is required on older rpmbuild (like on CentOS 5)
# fpm passes '--define buildroot ...' on the commandline, so just reuse that.
BuildRoot: %buildroot
# Add prefix, must not end with /
Prefix: /myapp 

Group: default
License: Private        
Vendor: foo 
URL: https://foobar.com            
Packager: fo@bar.com

%description
blah

%prep
# noop

%build
# noop

%install
# noop

%clean
# noop

%post

%preun

%files
%defattr(-,obama,madhouse,-)



%dir %attr(0755, obama, madhouse) 
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/bin
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/bin/win32
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/bin/win64
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/conf
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/data
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/docs
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/amqp
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/amqp/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/amqp/java/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/amqp/java/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/amqp/java/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/amqp/java/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/amqp/python
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/conf
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/java/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/java/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/java/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/java/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/websocket
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/websocket/css
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/websocket/img
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/websocket/js
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-composite-destinations
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-composite-destinations/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-composite-destinations/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-composite-destinations/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-composite-destinations/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-composite-destinations/src/main/java/example/composite
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-composite-destinations/src/main/java/example/composite/dest
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-durable-sub
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-durable-sub/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-durable-sub/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-durable-sub/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-durable-sub/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-durable-sub/src/main/java/example/topic
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-durable-sub/src/main/java/example/topic/durable
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-exclusive-consumer
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-exclusive-consumer/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-exclusive-consumer/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-exclusive-consumer/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-exclusive-consumer/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-exclusive-consumer/src/main/java/example/queue
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-exclusive-consumer/src/main/java/example/queue/exclusive
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-message-browser
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-message-browser/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-message-browser/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-message-browser/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-message-browser/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-message-browser/src/main/java/example/browser
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue/src/main/java/example/queue
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue/src/main/resources
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue-selector
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue-selector/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue-selector/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue-selector/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue-selector/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue-selector/src/main/java/example/queue
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue-selector/src/main/java/example/queue/selector
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-temp-destinations
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-temp-destinations/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-temp-destinations/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-temp-destinations/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-temp-destinations/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-temp-destinations/src/main/java/example/tempdest
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-topic
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-topic/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-topic/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-topic/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-topic/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-topic/src/main/java/example/topic
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-topic/src/main/resources
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-transaction
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-transaction/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-transaction/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-transaction/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-transaction/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-transaction/src/main/java/example/transaction
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-wildcard-consumer
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-wildcard-consumer/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-wildcard-consumer/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-wildcard-consumer/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-wildcard-consumer/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-wildcard-consumer/src/main/java/example/wildcard
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/cpp
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/ActiveMQExamples
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/ActiveMQExamples/Listener
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/ActiveMQExamples/Publisher
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/ecommerce
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/ecommerce/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/QueueMonitor
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/QueueRoundTrip
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/RequestReply
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/SelectorTalk
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/Talk
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/TransactedTalk
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/Chat
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/DurableChat
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/HierarchicalChat
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/MessageMonitor
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/RequestReply
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/SelectorChat
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/TransactedChat
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/conf
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/java/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/java/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/java/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/java/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/other
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/other/perfharness
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/cpp
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/ActiveMQExamples
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/ActiveMQExamples/Listener
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/ActiveMQExamples/Publisher
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/java/src
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/java/src/main
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/java/src/main/java
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/java/src/main/java/example
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/perl
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/php
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stompest
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stompest/async
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stompest/sync
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stomppy
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/ruby
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/websocket
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/websocket/css
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/websocket/img
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/websocket/js
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/lib
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/lib/camel
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/lib/extra
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/lib/optional
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/lib/web
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/META-INF
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/config
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/controller
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/filter
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/handler
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/jspf
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/form
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/jms
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/decorators
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/mochi
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/plotkit
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/styles
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/test
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/xml
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/api
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/api/WEB-INF
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/META-INF
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/WEB-INF
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/WEB-INF/classes
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/WEB-INF/classes/org
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/WEB-INF/classes/org/apache
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/WEB-INF/classes/org/apache/activemq
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/WEB-INF/classes/org/apache/activemq/util
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/images
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps/styles
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps-demo
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/META-INF
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/WEB-INF
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/mochi
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/plotkit
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/mqtt
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/portfolio
%dir %attr(0755, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/styles
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/test
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/test/assets
%dir %attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/websocket

%attr(0600, obama, madhouse) /apache-activemq-5.13.0/LICENSE
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/NOTICE
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/README.txt
%attr(0755, obama, madhouse) /apache-activemq-5.13.0/activemq-all-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/activemq
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/activemq-admin.bat
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/activemq.bat
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/activemq.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/win32/InstallService.bat
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/win32/UninstallService.bat
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/win32/activemq.bat
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/win32/wrapper.conf
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/win32/wrapper.dll
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/win32/wrapper.exe
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/win64/InstallService.bat
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/win64/UninstallService.bat
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/win64/activemq.bat
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/win64/wrapper.conf
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/win64/wrapper.dll
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/bin/win64/wrapper.exe
%attr(0700, obama, madhouse) /apache-activemq-5.13.0/bin/wrapper.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/activemq.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/broker-localhost.cert
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/broker.ks
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/broker.ts
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/client.ks
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/client.ts
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/credentials-enc.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/credentials.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/groups.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/jetty-realm.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/jetty.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/jmx.access
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/jmx.password
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/log4j.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/logging.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/login.config
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/conf/users.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/data/activemq.log
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/docs/WebConsole-README.txt
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/docs/index.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/docs/user-guide.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/amqp/java/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/amqp/java/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/amqp/java/src/main/java/example/Listener.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/amqp/java/src/main/java/example/Publisher.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/amqp/python/listener.py
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/amqp/python/publisher.py
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-demo.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-dynamic-network-broker1.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-dynamic-network-broker2.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-jdbc-performance.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-jdbc.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-leveldb-replicating.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-mqtt.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-scalability.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-security.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-specjms.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-static-network-broker1.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-static-network-broker2.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-stomp.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq-throughput.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/activemq.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/camel.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/jetty-demo.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/log4j.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/resin-web.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/conf/web.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/java/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/java/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/java/src/main/java/example/Listener.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/java/src/main/java/example/Publisher.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/websocket/css/bootstrap.min.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/websocket/css/bootstrap.min.responsive.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/websocket/img/glyphicons-halflings-white.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/websocket/img/glyphicons-halflings.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/websocket/index.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/websocket/js/jquery-1.7.2.min.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/websocket/js/mqttws31.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/mqtt/websocket/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-composite-destinations/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-composite-destinations/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-composite-destinations/src/main/java/example/composite/dest/Consumer.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-composite-destinations/src/main/java/example/composite/dest/Producer.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-durable-sub/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-durable-sub/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-durable-sub/src/main/java/example/topic/durable/Publisher.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-durable-sub/src/main/java/example/topic/durable/Subscriber.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-exclusive-consumer/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-exclusive-consumer/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-exclusive-consumer/src/main/java/example/queue/exclusive/Consumer.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-exclusive-consumer/src/main/java/example/queue/exclusive/Producer.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-message-browser/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-message-browser/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-message-browser/src/main/java/example/browser/Browser.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-message-browser/src/main/java/example/browser/Producer.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue/src/main/java/example/queue/Consumer.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue/src/main/java/example/queue/Producer.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue/src/main/resources/log4j.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue-selector/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue-selector/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue-selector/src/main/java/example/queue/selector/Consumer.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-queue-selector/src/main/java/example/queue/selector/Producer.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-temp-destinations/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-temp-destinations/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-temp-destinations/src/main/java/example/tempdest/Consumer.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-temp-destinations/src/main/java/example/tempdest/ProducerRequestReply.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-topic/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-topic/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-topic/src/main/java/example/topic/Publisher.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-topic/src/main/java/example/topic/Subscriber.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-topic/src/main/resources/log4j.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-transaction/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-transaction/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-transaction/src/main/java/example/transaction/Client.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-wildcard-consumer/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-wildcard-consumer/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/jms-example-wildcard-consumer/src/main/java/example/wildcard/Client.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/advanced-scenarios/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/cpp/Listener.cpp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/cpp/Publisher.cpp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/cpp/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/ActiveMQExamples/ActiveMQExamples.sln
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/ActiveMQExamples/ActiveMQExamples.userprefs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/ActiveMQExamples/Listener/AssemblyInfo.cs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/ActiveMQExamples/Listener/Listener.cs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/ActiveMQExamples/Listener/Listener.csproj
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/ActiveMQExamples/Listener/NMSTracer.cs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/ActiveMQExamples/Publisher/AssemblyInfo.cs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/ActiveMQExamples/Publisher/NMSTracer.cs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/ActiveMQExamples/Publisher/Publisher.cs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/ActiveMQExamples/Publisher/Publisher.csproj
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/csharp/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/ecommerce/README.txt
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/ecommerce/build.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/ecommerce/src/Retailer.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/ecommerce/src/Supplier.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/ecommerce/src/TransactionsDemo.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/ecommerce/src/Vendor.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/QueueMonitor/QueueMonitor.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/QueueMonitor/QueueMonitor.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/QueueRoundTrip/QueueRoundTrip.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/RequestReply/Replier.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/RequestReply/Requestor.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/SelectorTalk/SelectorTalk.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/Talk/Talk.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/QueuePTPSamples/TransactedTalk/TransactedTalk.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/Chat/Chat.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/DurableChat/DurableChat.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/HierarchicalChat/HierarchicalChat.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/MessageMonitor/MessageMonitor.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/MessageMonitor/MessageMonitor.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/RequestReply/TopicReplier.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/RequestReply/TopicRequestor.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/SelectorChat/SelectorChat.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/TopicPubSubSamples/TransactedChat/TransactedChat.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/build.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/conf/log4j.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/exploring-jms/readme.txt
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/java/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/java/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/java/src/main/java/example/Listener.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/java/src/main/java/example/Publisher.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/build.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/src/CommandLineSupport.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/src/ConsumerTool.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/src/EmbeddedBroker.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/src/Log4jJMSAppenderExample.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/src/ProducerAndConsumerTool.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/src/ProducerTool.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/src/RequesterTool.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/src/StompExample.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/src/TopicListener.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/src/TopicPublisher.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/src/jndi.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/openwire/swissarmy/src/log4j-jms.properties
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/other/perfharness/perfharness-activemq.sh
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/cpp/Listener.cpp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/cpp/Publisher.cpp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/cpp/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/ActiveMQExamples/ActiveMQExamples.sln
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/ActiveMQExamples/ActiveMQExamples.userprefs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/ActiveMQExamples/Listener/AssemblyInfo.cs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/ActiveMQExamples/Listener/Listener.cs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/ActiveMQExamples/Listener/Listener.csproj
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/ActiveMQExamples/Listener/NMSTracer.cs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/ActiveMQExamples/Publisher/AssemblyInfo.cs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/ActiveMQExamples/Publisher/NMSTracer.cs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/ActiveMQExamples/Publisher/Publisher.cs
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/ActiveMQExamples/Publisher/Publisher.csproj
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/csharp/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/java/pom.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/java/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/java/src/main/java/example/Listener.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/java/src/main/java/example/Publisher.java
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/perl/listener
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/perl/publisher
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/perl/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/php/listener.php
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/php/publisher.php
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/php/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stompest/async/__init__.py
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stompest/async/listener.py
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stompest/async/publisher.py
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stompest/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stompest/sync/__init__.py
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stompest/sync/listener.py
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stompest/sync/publisher.py
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stomppy/listener.py
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stomppy/publisher.py
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/python/stomppy/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/ruby/catstomp.rb
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/ruby/listener.rb
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/ruby/publisher.rb
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/ruby/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/ruby/stompcat.rb
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/websocket/css/bootstrap.min.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/websocket/css/bootstrap.min.responsive.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/websocket/img/glyphicons-halflings-white.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/websocket/img/glyphicons-halflings.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/websocket/index.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/websocket/js/jquery-1.7.2.min.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/websocket/js/stomp.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/examples/stomp/websocket/readme.md
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/activemq-broker-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/activemq-client-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/activemq-console-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/activemq-jaas-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/activemq-kahadb-store-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/activemq-openwire-legacy-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/activemq-protobuf-1.1.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/activemq-rar.txt
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/activemq-spring-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/activemq-web-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/camel/activemq-camel-5.13.0.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/camel/camel-core-2.16.1.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/camel/camel-jms-2.16.1.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/camel/camel-spring-2.16.1.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/extra/mqtt-client-1.12.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/geronimo-j2ee-management_1.1_spec-1.0.1.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/geronimo-jms_1.1_spec-1.1.1.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/geronimo-jta_1.0.1B_spec-1.0.1.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/hawtbuf-1.11.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/jcl-over-slf4j-1.7.13.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/activeio-core-3.1.4.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/activemq-amqp-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/activemq-http-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/activemq-jdbc-store-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/activemq-jms-pool-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/activemq-leveldb-store-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/activemq-log4j-appender-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/activemq-mqtt-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/activemq-partition-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/activemq-pool-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/activemq-runtime-config-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/activemq-shiro-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/activemq-stomp-5.13.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/commons-beanutils-1.8.3.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/commons-codec-1.9.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/commons-collections-3.2.2.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/commons-dbcp2-2.1.1.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/commons-lang-2.6.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/commons-net-3.3.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/commons-pool2-2.4.2.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/geronimo-j2ee-connector_1.5_spec-2.0.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/guava-12.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/hawtbuf-proto-1.11.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/hawtdispatch-1.22.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/hawtdispatch-scala-2.11-1.22.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/hawtdispatch-transport-1.22.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/hawtjni-runtime-1.9.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/httpclient-4.5.1.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/httpcore-4.4.4.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/insight-log-core-1.2.0.Beta4.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/insight-log4j-1.2.0.Beta4.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/jackson-annotations-2.6.3.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/jackson-core-2.6.3.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/jackson-databind-2.6.3.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/jasypt-1.9.2.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/jasypt-spring31-1.9.2.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/jaxb2-basics-runtime-0.6.4.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/jettison-1.3.7.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/jmdns-3.4.1.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/leveldb-0.6.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/leveldb-api-0.6.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/leveldbjni-1.8.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/log4j-1.2.17.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/org.apache.servicemix.bundles.josql-1.5_5.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/org.linkedin.util-core-1.4.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/org.linkedin.zookeeper-impl-1.4.0.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/proton-j-0.11.0.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/scala-library-2.11.0.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/shiro-core-1.2.4.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/shiro-spring-1.2.4.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/slf4j-log4j12-1.7.13.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/snappy-0.2.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/snappy-java-1.1.2.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/spring-aop-4.1.8.RELEASE.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/spring-beans-4.1.8.RELEASE.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/spring-context-4.1.8.RELEASE.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/spring-core-4.1.8.RELEASE.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/spring-expression-4.1.8.RELEASE.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/spring-jms-4.1.8.RELEASE.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/spring-oxm-4.1.8.RELEASE.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/spring-tx-4.1.8.RELEASE.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/velocity-1.7.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/xbean-spring-3.18.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/xpp3-1.1.4c.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/optional/xstream-1.4.8.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/optional/zookeeper-3.4.6.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/slf4j-api-1.7.13.jar
%attr(0640, obama, madhouse) /apache-activemq-5.13.0/lib/web/apache-el-8.0.9.M3.jar
%attr(0640, obama, madhouse) /apache-activemq-5.13.0/lib/web/apache-jsp-8.0.9.M3.jar
%attr(0640, obama, madhouse) /apache-activemq-5.13.0/lib/web/apache-jsp-9.2.13.v20150730.jar
%attr(0640, obama, madhouse) /apache-activemq-5.13.0/lib/web/apache-jstl-9.2.13.v20150730.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/web/asm-5.0.4.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/web/geronimo-annotation_1.0_spec-1.1.1.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/web/jdom-1.0.jar
%attr(0640, obama, madhouse) /apache-activemq-5.13.0/lib/web/jetty-all-9.2.13.v20150730.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/web/jolokia-core-1.3.2.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/web/json-simple-1.1.1.jar
%attr(0640, obama, madhouse) /apache-activemq-5.13.0/lib/web/org.eclipse.jdt.core-3.8.2.v20130121.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/lib/web/rome-1.0.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/web/spring-web-4.1.8.RELEASE.jar
%attr(0644, obama, madhouse) /apache-activemq-5.13.0/lib/web/spring-webmvc-4.1.8.RELEASE.jar
%attr(0640, obama, madhouse) /apache-activemq-5.13.0/lib/web/taglibs-standard-impl-1.2.1.jar
%attr(0640, obama, madhouse) /apache-activemq-5.13.0/lib/web/taglibs-standard-spec-1.2.1.jar
%attr(0640, obama, madhouse) /apache-activemq-5.13.0/lib/web/tomcat-servlet-api-8.0.24.jar
%attr(0640, obama, madhouse) /apache-activemq-5.13.0/lib/web/tomcat-websocket-api-8.0.24.jar
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/403.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/404.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/500.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/META-INF/LICENSE
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/META-INF/NOTICE
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/WebConsoleStarter$OsgiUtil.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/WebConsoleStarter.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/config/OsgiConfiguration.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/controller/CopyMessage.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/controller/CreateDestination.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/controller/CreateSubscriber.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/controller/DeleteDestination.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/controller/DeleteJob.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/controller/DeleteMessage.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/controller/DeleteSubscriber.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/controller/MoveMessage.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/controller/PurgeDestination.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/controller/SendMessage.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/filter/ApplicationContextFilter$1.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/filter/ApplicationContextFilter$2.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/filter/ApplicationContextFilter.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/classes/org/apache/activemq/web/handler/BindingBeanNameUrlHandlerMapping.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/dispatcher-servlet.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/jspf/headertags.jspf
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/form/checkbox.tag
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/form/escape.tag
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/form/forEachMapEntry.tag
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/form/option.tag
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/form/short.tag
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/form/text.tag
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/form/tooltip.tag
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/form/uri.tag
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/jms/forEachConnection.tag
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/jms/forEachMessage.tag
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/jms/formatTimestamp.tag
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/tags/jms/persistent.tag
%attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/web.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/webconsole-default.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/webconsole-embedded.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/webconsole-invm.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/webconsole-jndi.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/webconsole-osgi.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/webconsole-properties.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/WEB-INF/webconsole-query.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/browse.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/connection.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/connections.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/decorators/footer.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/decorators/head.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/decorators/header.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/graph.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/activemq-logo.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/asf-logo.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/big-bullet.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/black-footer-bottom.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/black-footer-left.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/black-footer-right.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/bottom-red-bar.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/checker-bg.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/content-left.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/content-right.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/feed_atom.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/feed_rss.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/left-box-bottom.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/left-box-right.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/left-box-top.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/oval-arrow.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/right-box-bottom.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/right-box-left.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/right-box-top.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/small-bullet-gray.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/small-bullet-red.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/spacer.gif
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/top-red-bar.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/white-header-left.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/white-header-right.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/images/white-header-top.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/index.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/common.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/css.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/mochi/MochiKit.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/mochi/__package__.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/plotkit/Base.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/plotkit/Canvas.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/plotkit/Layout.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/plotkit/SVG.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/plotkit/SweetCanvas.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/plotkit/SweetSVG.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/plotkit/dummy.svg
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/plotkit/iecanvas.htc
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/prettify.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/js/standardista-table-sorting.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/login.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/message.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/network.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/queueConsumers.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/queueGraph.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/queueProducers.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/queues.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/scheduled.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/send.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/slave.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/styles/prettify.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/styles/site.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/styles/sorttable.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/styles/type-settings.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/subscribers.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/test/dummy.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/test/index.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/test/systemProperties.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/topicProducers.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/topicSubscribers.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/topics.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/xml/queues.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/xml/subscribers.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/admin/xml/topics.jsp
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/api/WEB-INF/web.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/favicon.ico
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/META-INF/LICENSE
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/META-INF/NOTICE
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/WEB-INF/classes/org/apache/activemq/util/FilenameGuardFilter$GuardedHttpServletRequest.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/WEB-INF/classes/org/apache/activemq/util/FilenameGuardFilter.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/WEB-INF/classes/org/apache/activemq/util/IOHelper.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/WEB-INF/classes/org/apache/activemq/util/RestFilter.class
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/WEB-INF/web.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/fileserver/index.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/activemq-logo.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/asf-logo.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/big-bullet.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/black-footer-bottom.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/black-footer-left.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/black-footer-right.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/bottom-red-bar.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/checker-bg.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/content-left.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/content-right.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/feed_atom.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/feed_rss.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/left-box-bottom.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/left-box-right.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/left-box-top.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/oval-arrow.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/right-box-bottom.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/right-box-left.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/right-box-top.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/small-bullet-gray.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/small-bullet-red.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/spacer.gif
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/top-red-bar.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/white-header-left.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/white-header-right.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/images/white-header-top.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/index.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/styles/prettify.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/styles/site.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/styles/sorttable.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps/styles/type-settings.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/META-INF/LICENSE
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/META-INF/NOTICE
%attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/WEB-INF/web.xml
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/chat.css
%attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/chat.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/activemq-logo.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/asf-logo.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/big-bullet.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/black-footer-bottom.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/black-footer-left.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/black-footer-right.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/bottom-red-bar.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/checker-bg.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/content-left.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/content-right.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/feed_atom.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/feed_rss.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/left-box-bottom.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/left-box-right.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/left-box-top.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/oval-arrow.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/right-box-bottom.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/right-box-left.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/right-box-top.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/small-bullet-gray.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/small-bullet-red.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/spacer.gif
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/top-red-bar.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/white-header-left.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/white-header-right.png
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/images/white-header-top.png
%attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/index.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/amq.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/amq_dojo_adapter.js
%attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/amq_jquery_adapter.js
%attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/amq_prototype_adapter.js
%attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/chat.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/common.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/css.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/dojo.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/jquery-1.4.2.min.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/mochi/MochiKit.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/mochi/__package__.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/plotkit/Base.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/plotkit/Canvas.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/plotkit/Layout.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/plotkit/SVG.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/plotkit/SweetCanvas.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/plotkit/SweetSVG.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/plotkit/dummy.svg
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/plotkit/iecanvas.htc
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/prettify.js
%attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/prototype.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/js/standardista-table-sorting.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/mqtt/chat.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/mqtt/chat.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/mqtt/index.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/mqtt/mqttws31.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/portfolio/portfolio.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/portfolio/portfolio.js
%attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/send.html
%attr(0700, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/style.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/styles/prettify.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/styles/site.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/styles/sorttable.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/styles/type-settings.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/test/amq_test.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/test/assets/README
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/test/assets/jsunittest.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/test/assets/unittest.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/websocket/chat.css
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/websocket/chat.js
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/websocket/index.html
%attr(0600, obama, madhouse) /apache-activemq-5.13.0/webapps-demo/demo/websocket/stomp.js

%changelog
