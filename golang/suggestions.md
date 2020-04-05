## Non-blocking `select`

Use it if you have some sort looped task which could be interrupted by channel:

~~~go
for {
    select {
        case _, _ = <- messageBus:
            // Do smth
        default:
            // Exit to loop body
    }
}
~~~

## Use context for long-running IO procedures or for DB:

~~~
func ListenBroadcastUDP(ctx context.Context, packets chan []byte) error {
    // ...
}
~~~

~~~
type UserRepository interface {
    GetAllUsers(ctx context.Context, limit, offset uint64, options ...Option) ([]*model.User, error)
    // ...
}
~~~

## Dave Cheney's [errors](https://github.com/pkg/errors/) package gotcha

Do not wrap empty error with errors.Wrap() or you will lose it:

~~~
func Wrap(err error, message string) error {
    if err == nil {
        return nil
    }
    // ...
}
~~~

## Log error only once

It's easier to read error log if one error takes only a single line. Do not hesitate to use errors.Wrap() error to get a full track of error.

## Check is a channel not closed

There is a safe way to check is channel is not closed during continues read in a loop:

~~~
var (
    isChannelAlive bool
    message string
)

messageReadLoop:
for {
    select {
        case message, isChannelAlive = <- messageBus:
            if !isChannelAlive {
                break messageReadLoop
            }
            // regular processing    
        }
    }    
}
~~~

Some more details about it from here - https://www.youtube.com/watch?v=t9bEg2A4jsw

