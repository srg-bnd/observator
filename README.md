# Observator

Metrics collection and alerting service

## For students

### Updating the template

To be able to receive updates to autotests and other parts of the template, run the command:

```bash
git remote add -m main template https://github.com/Yandex-Practicum/go-musthave-metrics-tpl.git
```

To update the autotest code, run the command:

```bash
git fetch template && git checkout template/main .github
```

Then add the received changes to your repository.

### Launching autotests

To successfully run autotests, name the branches `iter<number>`, where `<number>` is the sequence number of the increment. For example, in the branch named `iter4`, autotests will be launched for increments from the first to the fourth.

When you merge a branch with an increment into the main branch `main`, all autotests will be run.

Read more about local and automatic startup in [README автотестов](https://github.com/Yandex-Practicum/go-autotests).
