// +build android

package main

/*
#cgo CFLAGS: -D__ANDROID_API__=16 --sysroot=/home/user/android-ndk-r14b/sysroot -isystem /home/user/android-ndk-r14b/sysroot/usr/include/arm-linux-androideabi -isystem /home/user/android-ndk-r14b/sources/cxx-stl/gnu-libstdc++/4.9/include -isystem /home/user/android-ndk-r14b/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a/include -fstack-protector-strong -DANDROID -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -fno-builtin-memmove -Os -mthumb -Wall -W -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -D__ANDROID_API__=16 --sysroot=/home/user/android-ndk-r14b/sysroot -isystem /home/user/android-ndk-r14b/sysroot/usr/include/arm-linux-androideabi -isystem /home/user/android-ndk-r14b/sources/cxx-stl/gnu-libstdc++/4.9/include -isystem /home/user/android-ndk-r14b/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a/include -fstack-protector-strong -DANDROID -march=armv7-a -mfloat-abi=softfp -mfpu=vfp -fno-builtin-memmove -O2 -Os -mthumb -std=gnu++11 -Wall -W -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../amlwwalker -I. -I/opt/Qt/5.10.0/android_armv7/include -I/opt/Qt/5.10.0/android_armv7/include/QtGui -I/opt/Qt/5.10.0/android_armv7/include/QtCore -I. -I/opt/Qt/5.10.0/android_armv7/mkspecs/android-g++
#cgo LDFLAGS: --sysroot=/home/user/android-ndk-r14b/platforms/android-16/arch-arm/  -Wl,-rpath=/opt/Qt/5.10.0/android_armv7/lib -Wl,--no-undefined -Wl,-z,noexecstack 
#cgo LDFLAGS:  -L/home/user/android-ndk-r14b/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L/home/user/android-ndk-r14b/toolchains/arm-linux-androideabi-4.9/prebuilt/linux-x86_64/bin/../lib/gcc/arm-linux-androideabi/4.9.x -L/opt/Qt/5.10.0/android_armv7/lib -lQt5Gui -L/opt/android/android-ndk-r10e/sources/cxx-stl/gnu-libstdc++/4.9/libs/armeabi-v7a -L/opt/android/android-ndk-r10e/toolchains/arm-linux-androideabi-4.9/prebuilt/linux-x86_64/bin/../lib/gcc/arm-linux-androideabi/4.9 -lQt5Core -lGLESv2 -lgnustl_shared -lgcc -llog -lz -lm -ldl -lc
#cgo LDFLAGS: -Wl,--allow-shlib-undefined
*/
import "C"
