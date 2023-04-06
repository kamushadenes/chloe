#!/usr/bin/env python3

import os
from datetime import datetime
from functools import reduce
from typing import AnyStr, Tuple, Dict, List

import openai
import requests


def list_releases(repository: AnyStr) -> List[Dict]:
    """List the releases from the GitHub API."""

    print('[*] Listing releases...')

    return requests.get(
        "https://api.github.com/repos/{}/releases".format(repository),
        headers={
            "Accept": "application/vnd.github+json",
            "Authorization": "Bearer {}".format(os.environ["GITHUB_TOKEN"])
        },
    ).json()


def sort_releases(releases: List[Dict]) -> List[Dict]:
    """Sort the releases by date."""

    print('[*] Sorting releases...')

    return sorted(releases,
                  key=lambda d: datetime.strptime(d['created_at'], "%Y-%m-%dT%H:%M:%SZ"),
                  reverse=True)


def get_latest_draft(releases: List[Dict]) -> Dict:
    """Get the release from the GitHub API."""

    print('[*] Getting latest draft release...')

    for release in releases:
        if release['draft']:
            return release


def get_improved_release_message(release: Dict) -> Tuple[Dict, AnyStr]:
    """Improves the release message using OpenAI's GPT-3 API."""

    print('[*] Improving release message...')

    completion = openai.ChatCompletion.create(model='gpt-3.5-turbo', messages=[
        {
            "role": "system",
            "content": "Your task is to rewrite release notes in a more concise manner, "
                       "no need to mention specific commits. "
                       "Group things by features / bug fixes / etc as appropriate. "
                       "Try to focus on the most important changes. "
                       "Return it in markdown format.",
        },
        {
            "role": "user",
            "content": release['body']
        }
    ])

    return release, completion.choices[0].message.content


def update_release_notes(args: Tuple[Dict, AnyStr]) -> None:
    """Update the release notes using the GitHub API."""

    print('[*] Updating release notes...')

    print('[*] New content:\n\n{}\n\n'.format(args[1]))

    r = requests.patch(
        args[0]['url'],
        headers={
            "Accept": "application/vnd.github+json",
            "Authorization": "Bearer {}".format(os.environ["GITHUB_TOKEN"])
        },
        json={
            "body": args[1],
            "draft": False,
        }
    )

    if r.status_code != 200:
        print("[-] Failed to update release notes: {}".format(r.text))
    else:
        print('[+] Successfully updated release notes!')


if __name__ == '__main__':
    if "GITHUB_TOKEN" not in os.environ:
        raise ValueError("GITHUB_TOKEN environment variable is not set.")

    if "OPENAI_API_KEY" not in os.environ:
        raise ValueError("OPENAI_API_KEY environment variable is not set.")

    reduce(lambda x, f: f(x),
           [
               list_releases,
               sort_releases,
               get_latest_draft,
               get_improved_release_message,
               update_release_notes,
           ], "kamushadenes/chloe")
