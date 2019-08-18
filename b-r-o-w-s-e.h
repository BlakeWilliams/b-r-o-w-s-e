#ifndef BROWSE_H
#define BROWSE_H

#import <Cocoa/Cocoa.h>

extern void HandleURL(char*);

@interface BrowseAppDelegate: NSObject<NSApplicationDelegate>
  - (void)handleGetURLEvent:(NSAppleEventDescriptor *) event withReplyEvent:(NSAppleEventDescriptor *)replyEvent;
@end

void RunApp(void);

#endif
