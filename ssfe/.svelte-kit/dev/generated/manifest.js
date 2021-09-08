const c = [
	() => import("../../../src/routes/__layout.svelte"),
	() => import("../components/error.svelte"),
	() => import("../../../src/routes/index.svelte"),
	() => import("../../../src/routes/asset.svelte"),
	() => import("../../../src/routes/orgs/__layout.svelte"),
	() => import("../../../src/routes/orgs/index.svelte"),
	() => import("../../../src/routes/orgs/[id]/index.svelte"),
	() => import("../../../src/routes/orgs/[id]/assets.svelte"),
	() => import("../../../src/routes/org.svelte")
];

const d = decodeURIComponent;

export const routes = [
	// src/routes/index.svelte
	[/^\/$/, [c[0], c[2]], [c[1]]],

	// src/routes/asset.svelte
	[/^\/asset\/?$/, [c[0], c[3]], [c[1]]],

	// src/routes/orgs/index.svelte
	[/^\/orgs\/?$/, [c[0], c[4], c[5]], [c[1]]],

	// src/routes/orgs/[id]/index.svelte
	[/^\/orgs\/([^/]+?)\/?$/, [c[0], c[4], c[6]], [c[1]], (m) => ({ id: d(m[1])})],

	// src/routes/orgs/[id]/assets.svelte
	[/^\/orgs\/([^/]+?)\/assets\/?$/, [c[0], c[4], c[7]], [c[1]], (m) => ({ id: d(m[1])})],

	// src/routes/org.svelte
	[/^\/org\/?$/, [c[0], c[8]], [c[1]]]
];

// we import the root layout/error components eagerly, so that
// connectivity errors after initialisation don't nuke the app
export const fallback = [c[0](), c[1]()];