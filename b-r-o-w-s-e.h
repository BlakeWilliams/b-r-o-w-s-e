#import <Cocoa/Cocoa.h>

extern void HandleURL(char*);

@interface BrowseAppDelegate: NSObject<NSApplicationDelegate>
  - (void)handleGetURLEvent:(NSAppleEventDescriptor *) event withReplyEvent:(NSAppleEventDescriptor *)replyEvent;
@end

void RunApp();
void ShowAlert(char*, char*);
