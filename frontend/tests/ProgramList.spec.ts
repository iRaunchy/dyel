
import { describe, it, expect, vi, beforeEach } from "vitest";
import { mount } from "@vue/test-utils";
import ProgramList from "../src/components/ProgramList.vue"; // Adjust path if needed
import { NButton, NAlert, NAvatar, NEmpty, NPageHeader, NSpace, NSpin } from "naive-ui";

// Mock the utility functions
vi.mock("../src/utils/utils", () => ({
    formatDate: vi.fn(() => 'Formatted Date'),
    getProgramColor: vi.fn(() => '#test-color'),
    getInitials: vi.fn(() => 'TP')
}));

// Mock vue-router
vi.mock("vue-router", () => ({
    useRouter: () => ({
        push: vi.fn()
    })
}));

describe("ProgramList", () => {
    beforeEach(() => {
        vi.resetAllMocks();

        // Mock fetch globally
        global.fetch = vi.fn();
    });

    it('shows loading state initially', async () => {
        // Setup fetch to never resolve during this test
        vi.mocked(global.fetch).mockReturnValue(new Promise(() => {}));

        const wrapper = mount(ProgramList, {
            global: {
                stubs: {
                    // Stub all Naive UI components with true (simple placeholder element)
                    'n-page-header': true,
                    'n-space': true,
                    'n-spin': true,
                    'n-empty': true,
                    'n-alert': true,
                    'n-avatar': true,
                    'n-button': true,
                },
            }
        });

        // Verify loading state - using direct access to the component's state
        expect(wrapper.vm.loading).toBe(true);
    });

    it('shows empty state when no programs', async () => {
        // Mock successful fetch with empty array
        vi.mocked(global.fetch).mockResolvedValue({
            ok: true,
            json: () => Promise.resolve([])
        } as unknown as Response);

        const wrapper = mount(ProgramList, {
            global: {
                stubs: {
                    // Stub all Naive UI components with true (simple placeholder element)
                    'n-page-header': true,
                    'n-space': true,
                    'n-spin': true,
                    'n-empty': true,
                    'n-alert': true,
                    'n-avatar': true,
                    'n-button': true,
                },
            }
        });

        // Wait for component to update after fetch completion
        await vi.waitFor(() => {
            expect(wrapper.vm.loading).toBe(false);
        });

        // Check for empty programs array
        expect(wrapper.vm.programs.length).toBe(0);

        // Check for empty state text
        expect(wrapper.text()).toContain('No programs yet');
    });
});