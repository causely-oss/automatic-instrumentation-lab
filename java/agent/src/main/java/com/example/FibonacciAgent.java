package com.example;

import java.lang.instrument.Instrumentation;
import java.util.Stack;
import java.util.concurrent.Callable;
import java.util.concurrent.atomic.AtomicLong;
import net.bytebuddy.agent.builder.AgentBuilder;
import net.bytebuddy.implementation.MethodDelegation;
import net.bytebuddy.implementation.bind.annotation.*;
import net.bytebuddy.matcher.ElementMatchers;

public class FibonacciAgent {
    private static final ThreadLocal<Stack<CallInfo>> callStack = ThreadLocal.withInitial(Stack::new);
    private static final AtomicLong callCounter = new AtomicLong(0);

    static class CallInfo {
        final int n;
        final long startTime;
        final long callId;
        
        CallInfo(int n, long startTime, long callId) {
            this.n = n;
            this.startTime = startTime;
            this.callId = callId;
        }
    }

    public static void premain(String args, Instrumentation inst) {
        System.out.println("Fibonacci Agent loaded!");
        
        new AgentBuilder.Default()
            .type(ElementMatchers.nameStartsWith("com.example.FibonacciApp"))
            .transform((builder, typeDescription, classLoader, module, protectionDomain) ->
                builder.method(ElementMatchers.named("fibonacci"))
                       .intercept(MethodDelegation.to(FibonacciInterceptor.class))
            ).installOn(inst);
    }

    public static class FibonacciInterceptor {
        
        @RuntimeType
        public static Object intercept(@Origin String methodName,
                                    @AllArguments Object[] args,
                                    @SuperCall Callable<?> callable) throws Exception {
            
            int n = (Integer) args[0];
            
            // Start timing
            long callId = callCounter.incrementAndGet();
            long startTime = System.nanoTime();
            CallInfo callInfo = new CallInfo(n, startTime, callId);
            callStack.get().push(callInfo);
            
            try {
                // Call the original method
                Object result = callable.call();
                
                // End timing
                Stack<CallInfo> stack = callStack.get();
                if (!stack.isEmpty()) {
                    CallInfo completedCall = stack.pop();
                    long duration = System.nanoTime() - completedCall.startTime;
                    System.out.printf("Call %d: fibonacci(%d) took %d ns%n", 
                                    completedCall.callId, completedCall.n, duration);
                }
                
                return result;
            } catch (Exception e) {
                // Clean up stack on exception
                Stack<CallInfo> stack = callStack.get();
                if (!stack.isEmpty()) {
                    stack.pop();
                }
                throw e;
            }
        }
    }
}
