package time_util

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const UserTimezoneKey = "tz_info"

// TimeZoneInterceptor Extracts metadata and inserts into context
func TimeZoneInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		// Extract user timezone
		if tz, exists := md[UserTimezoneKey]; exists && len(tz) > 0 {
			ctx = context.WithValue(ctx, UserTimezoneKey, tz[0])
		} else {
			ctx = context.WithValue(ctx, UserTimezoneKey, "UTC") // Default to UTC
		}
	}

	// Call the next handler with the modified context
	return handler(ctx, req)
}
