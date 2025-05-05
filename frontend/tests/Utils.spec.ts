import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import { formatDate, getInitials, getProgramColor } from "../src/utils/utils";

describe("utils", () => {
    describe("formatDate", () => {
        beforeEach(() => {
            // Mock current date to make tests deterministic
            vi.useFakeTimers()
            vi.setSystemTime(new Date(2023, 0, 15)) // January 15, 2023
        })

        afterEach(() => {
            vi.useRealTimers()
        })

        it('returns empty string for empty input', () => {
            expect(formatDate('')).toBe('')
        })

        it('returns "Today" for today', () => {
            const today = new Date(2023, 0, 15).toISOString()
            expect(formatDate(today)).toBe('Today')
        })

        it('returns "Yesterday" for yesterday', () => {
            const yesterday = new Date(2023, 0, 14).toISOString()
            expect(formatDate(yesterday)).toBe('Yesterday')
        })
    })

    describe("getInitials", () => {
        it('returns uppercase initials for single word', () => {
            expect(getInitials('Cardio')).toBe('C')
        })

        it('returns first two uppercase initials for multi-word name', () => {
            expect(getInitials('Full Body Workout')).toBe('FB')
        })
    })

    describe("getProgramColor", () => {
        it('returns consistent color for same name', () => {
            const color1 = getProgramColor('Test Program')
            const color2 = getProgramColor('Test Program')
            expect(color1).toBe(color2)
        })
    })
})