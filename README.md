# dashboard
Dashboard for etcd. You can still use the cool etcd dashboard with etcd 2.*.

## Getting dashboard

The latest release and setup instructions are available at [Github](https://github.com/AdoHe/etcd-dashboard).

## Building

You can build dashboard from source:

```
git clone git@github.com:AdoHe/etcd-dashboard.git
cd etcd-dashboard
./build
```
This will generate binary called `./bin/etcd-dashboard`.

## Running

Just run this command:

```
./bin/etcd-dashboard -host 127.0.0.1 -port 8080
```
And then type this `http://127.0.0.1:8080/mod/dashboard` address in your favorite browser, you will see the awesome etcd dashboard.**Just try it**
