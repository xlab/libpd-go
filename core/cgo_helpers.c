// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sat, 27 Aug 2016 22:58:56 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

#include "_cgo_export.h"
#include "cgo_helpers.h"

void t_libpd_printhook_e41f15cd(char* recv) {
	printHookE41F15CD(recv);
}

void t_libpd_banghook_7ee06eb6(char* recv) {
	bangHook7EE06EB6(recv);
}

void t_libpd_floathook_c0728ac0(char* recv, float x) {
	floatHookC0728AC0(recv, x);
}

void t_libpd_symbolhook_a2f11e30(char* recv, char* sym) {
	symbolHookA2F11E30(recv, sym);
}

void t_libpd_listhook_c08dcef4(char* recv, int argc, t_atom* argv) {
	listHookC08DCEF4(recv, argc, argv);
}

void t_libpd_messagehook_cb0280(char* recv, char* msg, int argc, t_atom* argv) {
	messageHookCB0280(recv, msg, argc, argv);
}

void t_libpd_noteonhook_73e8687d(int channel, int pitch, int velocity) {
	noteOnHook73E8687D(channel, pitch, velocity);
}

void t_libpd_controlchangehook_7248f877(int channel, int controller, int value) {
	controlChangeHook7248F877(channel, controller, value);
}

void t_libpd_programchangehook_527988c8(int channel, int value) {
	programChangeHook527988C8(channel, value);
}

void t_libpd_pitchbendhook_8bf1c1f3(int channel, int value) {
	pitchbendHook8BF1C1F3(channel, value);
}

void t_libpd_aftertouchhook_e7d4f475(int channel, int value) {
	aftertouchHookE7D4F475(channel, value);
}

void t_libpd_polyaftertouchhook_c3cfeaf7(int channel, int pitch, int value) {
	polyAftertouchHookC3CFEAF7(channel, pitch, value);
}

void t_libpd_midibytehook_6393d784(int port, int byte) {
	mIDIByteHook6393D784(port, byte);
}

