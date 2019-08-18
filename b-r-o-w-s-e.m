#include "b-r-o-w-s-e.h"

@implementation BrowseAppDelegate
- (void)applicationWillFinishLaunching:(NSNotification *)aNotification
{
  NSAppleEventManager *appleEventManager = [NSAppleEventManager sharedAppleEventManager];
  [appleEventManager setEventHandler:self
                         andSelector:@selector(handleGetURLEvent:withReplyEvent:)
                         forEventClass:kInternetEventClass andEventID:kAEGetURL];
}

- (NSApplicationTerminateReply)applicationShouldTerminate:(NSApplication *)sender
{
  return NSTerminateNow;
}

- (void)handleGetURLEvent:(NSAppleEventDescriptor *)event
           withReplyEvent:(NSAppleEventDescriptor *)replyEvent {
  HandleURL((char*)[[[event paramDescriptorForKeyword:keyDirectObject] stringValue] UTF8String]);
}
@end

void RunApp(void) {
  [NSAutoreleasePool new];
  [NSApplication sharedApplication];
  BrowseAppDelegate *app = [BrowseAppDelegate alloc];
  [NSApp setDelegate:app];
  [NSApp run];
}

void ShowAlert(char* cmessage, char* cdetails) {
  NSString *message = [NSString stringWithCString:cmessage encoding:NSUTF8StringEncoding];
  NSString *details = [NSString stringWithCString:cdetails encoding:NSUTF8StringEncoding];

  NSAlert *alert = [[NSAlert alloc] init];
  [alert setMessageText:message];
  [alert setInformativeText:details];
  [alert addButtonWithTitle:@"Okay"];
  [alert runModal];
}
