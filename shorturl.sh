#!/bin/bash

name="Short URL Web Server"

APP_HOME=$(cd "$(dirname "$0")"; pwd)
APP_BIN=$APP_HOME/shorturl
PID_FILE=/tmp/shorturl.pid

start() {
    $APP_BIN > /dev/null &

    echo -e "\033[32m$name started.\033[0m"

    return 0
}

stop() {
    if status ; then
        pid=`cat "$PID_FILE"`
        echo -e "\033[33mKilling $name (pid $pid) with SIGTERM.\033[0m"
        kill -TERM $pid

        # Wait for it to exit.
        for i in 1 2 3 4 5 6 7 8 9 ; do
            echo -e "\033[33mWaiting $name (pid $pid) to die...\033[0m"
            status || break
            sleep 1
        done

        if status ; then
            echo -e "\033[31m$name stop failed; still running.\033[0m"
            return 1 # stop timed out and not forced
        else
            echo -e "\033[32m$name stopped.\033[0m"
        fi
    else
        echo -e "\033[31m$name is not running.\033[0m"
    fi
}

status() {
    if [ -f "$PID_FILE" ] ; then
        pid=`cat "$PID_FILE"`
        if kill -0 $pid > /dev/null 2> /dev/null ; then
            return 0
        else
            return 2 # program is dead but pid file exists
        fi
    else
        return 3 # program is not running
    fi
}

case "$1" in
    start)
        status
        code=$?
        if [ $code -eq 0 ]; then
            echo -e "\033[33m$name is already running.\033[0m"
        else
            start
            code=$?
        fi

        exit $code
        ;;

    stop)
        stop
        ;;

    status)
        status
        code=$?
        if [ $code -eq 0 ] ; then
            echo -e "\033[32m$name is running.\033[0m"
        else
            echo -e "\033[31m$name is not running.\033[0m"
        fi

        exit $code
        ;;

    restart)
        stop && start
        ;;
    *)

        echo "Usage: $0 {start|stop|status|restart}" >&2
        exit 3
    ;;
esac

exit $?
