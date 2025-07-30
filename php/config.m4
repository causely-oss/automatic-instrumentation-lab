dnl $Id$
dnl config.m4 for extension observer

PHP_ARG_ENABLE(observer, whether to enable observer support,
[  --enable-observer           Enable observer support])

if test "$PHP_OBSERVER" != "no"; then
  PHP_NEW_EXTENSION(observer, observer.c, $ext_shared)
fi 