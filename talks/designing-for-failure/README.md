# Designing for failure

## Abstract

In recent years, we’ve been talking more and more about failure and
resilience in a distributed system architecture. Unfortunately, failure has
been neglected for quite some time in our industry. Systems have grown to
be much more complex and challenging to deal with, especially in the
Kubernetes era. With all this new complexity it comes the question: how can
we design systems to be resilient and ready to fail?
This is something that mission-critical systems always had to think of first,
but many of our services (until today) neglect the importance of this. There
are many concepts we can learn and use from other fields like electronics,
aviation or naval industries that will help us be prepared for the unexpected.

In this talk, I would like to talk about how to design a system for failure.
What are the pitfalls and gotchas that we have to be prepared in a
microservice environment? How can SRE principles help us get there? And
most importantly, how do we put this into practice?

## Outline

In all of the steps I’ll show some Go, Kubernetes or Istio code on how these
patterns can be applied.

* Introduction
* Who am I?
* Learning from aviation industry
* Latency control
    * Timeouts
    * Circuit Breakers
    * Retries
* Load Shedding
    * Load Balancing
    * Outlier Server Host Detection
* Isolation
    * Bulkhead
    * Learning from naval industry
* Self Healing
    * Health Checks
    * Outbox Pattern
* Traffic Control
    * Rate Limiters
* Service Mashes
    * Istio
* Testing
    * Chaos Monkey
* Observability
    * Metrics
    * Distributed Tracing
    * Logging
    * Open Telemetry
* Conclusion
* Questions

## Presentations

1. Golab `2019`
- Slides: https://speakerdeck.com/italolelis/designing-for-failure
- Video: https://youtu.be/BOn3R41UrV8

2. Golang Piter `2019`
- Slides: https://speakerdeck.com/italolelis/designing-for-failure-4d969755-fbce-4b88-87ba-df8ea50d3dca
- Video: https://youtu.be/QWRPWb1Tzqs

2. Godays `2019`
- Slides: https://speakerdeck.com/italolelis/designing-for-failure-4d969755-fbce-4b88-87ba-df8ea50d3dca
- Video: https://youtu.be/DKhC_XH8fDs

