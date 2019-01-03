footer: Go Days Berlin 2019

## Desining for failure
### @italolelis

![](images/intro.jpg)

---

## Who __am__ I?

![](images/italo.jpg)

---

![](images/control_panel.jpg)

^ Think about an airplane, a boing 777 to be especific
^ It has something called ELMS. It controls all eletric circuits in the plane
^ If this system fails there are at least 4 fallbacks that are automatically activated to keep the essentials eletric systems running

---

## __Essentials__

^ The important word here is essentials
^ The plane keeps running without problems for a limited period of time and without manual intervention
^ This system was created on 1993

---

> Resilience is a Requirement, Not a Feature
-- Liang Guo

---

## Dependency Isolation and Graceful Degradation

^ In a distributed services architecture, a service will often have multiple downstream dependencies
^ When the number of dependencies increase, the probability of a single problematic dependency bringing down a service becomes higher.

---

## Health-check and Load Balancing

![](images/heartbeat.jpg)

---

![fit](images/health-code.png)

---

## If everything is OK you get...

---

![fit](images/health-ok.png)

---

## If things are not good but your app still can work...

---

![fit](images/health-partial.png)

---

## Otherwise...

---

![fit](images/health-fail.png)

---

## Self-healing

---

## Load shedding

^ Circuit Breakers
^ API Request Deadline
^ Request Queuing

---

## Retry Logic

^ CB are more efficient when they have a retry logic with exponential backoff

---

## Bulkhead

---

## Rate Limiters

^ Client Quota-based Rate Limit

---

## Failover caching

---

## Event Driven Architecture

---

## Outbox Pattern

^ How can producers reliably send messages when the broker/consumer is unavailable?
^ Lets imagine that you are pushing events into rabbitmq. Now think that for some reason rabbitmq is not available
^ What happens to the messages?

---

## SLO's and SLI's

---

## Monitoring

---

![fit](images/monitoring-code.png)

---

## Distributed Tracing

---

![fit](images/trace-code.png)

---

## Open Census

---

![fit](images/opencensus.png)

---

## *Questions?*

![](images/questions.jpg)

---

![](images/thanks.jpg)