const std = @import("std");
const testing = std.testing;

export const CalibrationValue = struct {
    data: []const u8,
    value: u32,

    pub fn calculate() u32 {
        return 0;
    }
};

export fn read_input(i: []const u8) []CalibrationValue {
    _ = i;}

test "basic add functionality" {
    try testing.expect(add(3, 7) == 10);
}
