#include "php.h"
#include "Zend/zend_observer.h"
#include <time.h>

// Declare the PHP_MINIT_FUNCTION before using it in the module entry
PHP_MINIT_FUNCTION(observer);

// Global variables to store start time and cumulative time
static clock_t start_time;
static double cumulative_time = 0.0;

// Observer function to log execution time of fib() function
static void observer_begin(zend_execute_data *execute_data) {
    if (execute_data->func && execute_data->func->common.function_name) {
        const char *function_name = ZSTR_VAL(execute_data->func->common.function_name);
        if (function_name && strcmp(function_name, "fib") == 0) {
            start_time = clock();
        }
    }
}

static void observer_end(zend_execute_data *execute_data, zval *retval) {
    if (execute_data->func && execute_data->func->common.function_name) {
        const char *function_name = ZSTR_VAL(execute_data->func->common.function_name);
        if (function_name && strcmp(function_name, "fib") == 0) {
            clock_t end_time = clock();
            double duration = (double)(end_time - start_time) / CLOCKS_PER_SEC * 1000;
            cumulative_time += duration;
            php_printf("Function %s() took %.2f ms, cumulative time: %.2f ms\n", function_name, duration, cumulative_time);
        }
    }
}

// Observer hook
static zend_observer_fcall_handlers observer_fcall_handlers(zend_execute_data *execute_data) {
    zend_observer_fcall_handlers handlers = {observer_begin, observer_end};
    return handlers;
}

// Module entry
zend_module_entry observer_module_entry = {
    STANDARD_MODULE_HEADER,
    "observer",
    NULL,
    PHP_MINIT(observer),
    NULL,
    NULL,
    NULL,
    NULL,
    NO_VERSION_YET,
    STANDARD_MODULE_PROPERTIES
};

ZEND_GET_MODULE(observer)

// Module startup
PHP_MINIT_FUNCTION(observer) {
    php_printf("Observer module loaded\n"); // Message when module is loaded
    zend_observer_fcall_register(observer_fcall_handlers);
    return SUCCESS;
}
