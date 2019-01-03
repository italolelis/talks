footer: Go Days Berlin 2019

## Desining for failure
### @italolelis

![](images/intro.jpg)

---

## Who __am__ I?

![](images/italo.jpg)

---

![](images/control_panel.jpg)

^ Think about an airplane
^ It has something called EMLS. It controls all eletric circuits in the plane
^ If this system fails there are at least 2 more fallbacks that are automatically activated that keep the essentials eletric systems running.

---

## __Essentials__

^ The important word here is essentials

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