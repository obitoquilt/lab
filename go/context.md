# context

```golang
var (
// An emptyCtx is never canceled, has no values, and has no deadline. It is not struct{}, since vars of this type must have distinct addresses.
type emptyCtx int

background = new(emptyCtx)
todo       = new(emptyCtx)
)
```

## Backgroud

它永远不会被取消，没有值，也没有截止日期。它通常由 main 函数、初始化和测试使用，并作为传入请求的顶级上下文。

## TODO 

当不清楚使用哪个上下文或它还不可用时（因为周围的函数还没有扩展到接受上下文参数）

## cancelCtx

```golang
type cancelCtx struct {
	Context                        // parent context

	mu       sync.Mutex            // protects following fields
	done     atomic.Value          // of chan struct{}, created lazily, closed by first cancel call
	children map[canceler]struct{} // set to nil by the first cancel call
	err      error                 // set to non-nil by the first cancel call
	cause    error                 // set to non-nil by the first cancel call
}
```