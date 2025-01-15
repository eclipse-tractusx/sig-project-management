---
name: Product Release Checks
about: Use this template to track all release checks for your component
title: "[Product] Release Checks"
---

<!-- 
Thanks for your contribution! Please fill out this template as good as possible. 
Important: Contributing Guidelines can be found here: https://eclipse-tractusx.github.io/docs/oss/how-to-contribute
Checkout the repository README for process description. 
-->

## Release Info

Please provide information on what you want to be included in the Eclipse Tractus-X release.
If you are not owner of this issue, please provide the information as comment to the issue.
Make sure to assign this issue to expert(s) for their approval, as soon as you have finished preparation. Multiple assignees allowed; they will un-assign themselves once review completed.

**Responsible contacts**:
<!-- For this new release -->
- Committer: *GitHub handle*
- ExpertGroup: *GitHub handle*
- Contributor contact: *GitHub handle*

**Version to be included in Eclipse Tractus-X release**:
<!-- Responsibility: Committer -->

- App version: *version placeholder*
- Helm Chart version: *version placeholder*

**Leading product repository:** *repository link*

## Compliance
<!-- Responsibility: Committer, ExpertGroup, Technical Committee for Standardization -->

- [ ] Possible changes (through the new features) on related Catena-X standards are considered and addressed
- [ ] [**Tractus-X Release Guidelines**](https://eclipse-tractusx.github.io/docs/release)(TRGs) fulfilled

Make sure to open and fill in a separate **documentation issue** in your product repository using the [Quality Gate Checklist issue template](https://github.com/eclipse-tractusx/.github/blob/main/.github/ISSUE_TEMPLATE/qg-checklist.md)

> [!NOTE]
> Note: most criteria for documentation and security are now covered in TRGs

## Functionality
<!-- Responsibility: Committer, Testmanagement -->

- [ ] Feature works as expected and described
- [ ] Backward compatibility maintained (depending Major/ Minor) 

## Performance
<!-- Responsibility: Testmanagement -->

- [ ] Memory and CPU usage is within acceptable limits oriented with goldilocks 

## Testing
<!-- Responsibility: Committer, Testmanagement -->

- [ ] Unit tests cover all new features
- [ ] Integration tests are updated
- [ ] E2E/Integration test passed
- [ ] Regressions tests passed

## Feature summary
<!-- 
Responsibility:
- Committer -> adding the features
- Testmanagement -> approval for teststatus
-->

Please provide a list of all features that have been developed in the current phase. This list is essential for several reasons:

- **Test Preparation:** Identifying new features will help in the creation of new end-to-end (E2E) tests and ensure that the relevant aspects are thoroughly tested.
- **Awareness:** It will raise awareness of the features that need to be prioritized during the testing process.
- **Board Management:** This will assist in keeping the project board organized and up-to-date.

By documenting these features, you will contribute to a smoother and more efficient release process.

> [!NOTE]
> Note: all features needs to be closed to get approval for release

| Feature | Test Status | Note |
|----------|----------|----------|
| linked feature 1 | Test status | note 1 |
| linked feature 2 | Test status | note 2 |
| linked feature 3 | Test status | note 3 |

## Standard summary

Please provide a list of all standards that have been affected. By documenting these standards, you will contribute to a smoother and more efficient release process.

- [ ] standard name

## Release Documentation
<!-- Responsibility: Committer, Testmanagement, Releasemanagement -->

- [ ] Release notes/Changelogs are updated
- [ ] Migration scripts/documentation are included if necessary (Ensure that any database or infrastructure migrations are included).
- [ ] Known knowns
  - topic 1
  - topic 2

## Summary
<!-- Responsibility: Committer -->

Please provide a short summary about the new values/benefits of the new features here:

- description of the value of feature 1
- description of the value of feature 2

## Helpful Links

- [Eclipse Tractus-X API Hub](https://eclipse-tractusx.github.io/api-hub)
